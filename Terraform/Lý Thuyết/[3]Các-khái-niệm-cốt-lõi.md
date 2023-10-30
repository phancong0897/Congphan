# Các khái niệm cốt lõi

### provider

- Một provider làm việc khá nhiều như thiết bị điều khiển một hệ điều hành của. Nó hiển thị một tập hợp các loại tài nguyên bằng 
cách sử dụng một sự trừu tượng chung, do đó che đi các chi tiết về cách tạo, sửa đổi và hủy một tài nguyên khá rõ ràng đối với người dùng .

- Terraform tự động tải xuống các nhà cung cấp từ public registry của nó khi cần, dựa trên các tài nguyên của một dự án nhất định. Nó cũng có thể sử dụng các plugin tùy chỉnh mà người dùng phải cài đặt thủ công.

- Với một vài ngoại lệ, việc sử dụng trình cung cấp yêu cầu phải định cấu hình trình cung cấp với một số tham số . Những điều đó khác nhau rất nhiều giữa các nhà cung cấp, nhưng nói chung, chúng ta sẽ cần cung cấp thông tin đăng nhập để nó có thể tiếp cận API của mình và gửi yêu cầu.

- Mặc dù không hoàn toàn cần thiết, nhưng việc khai báo rõ ràng nhà cung cấp nào mà chúng ta sẽ sử dụng trong project Terraform của mình và thông báo về phiên bản của nó được coi là một phương pháp hay. Vì mục đích này, chúng ta sử dụng thuộc tính version có sẵn cho bất kỳ khai báo provider nào :

    ```
    provider "kubernetes" {
      version = "~> 1.10"
    }

    ```

- Tại đây, vì chúng ta không cung cấp bất kỳ thông số bổ sung nào, Terraform sẽ tìm kiếm các thông số cần thiết ở nơi khác . Trong trường hợp này, việc triển khai của nhà cung cấp sẽ tìm kiếm các tham số kết nối bằng cách sử dụng cùng các vị trí được kubectl sử dụng . Các phương pháp phổ biến khác là sử dụng biến môi trường và tệp biến , chúng chỉ là tệp chứa các cặp khóa-giá trị.

### Resource

- Trong Terraform, resource là bất kỳ thứ gì có thể là mục tiêu cho các hoạt động CRUD trong ngữ cảnh của một provider nhất định. Một số ví dụ là phiên bản EC2, Azure MariaDB hoặc DNS.

- Hãy xem xét một định nghĩa tài nguyên đơn giản:
    ```

    resource "aws_instance" "web" {
      ami = "some-ami-id"
      instance_type = "t2.micro"
    }

    ```

- Đầu tiên, chúng ta luôn có từ khóa resource bắt đầu định nghĩa. Tiếp theo, chúng ta có loại resource , loại resource này thường tuân theo quy ước provider_type cấp . Trong ví dụ trên, aws_instance là một loại tài nguyên do nhà cung cấp AWS xác định, được sử dụng để xác định một phiên bản EC2. Sau đó, có tên tài nguyên do người dùng xác định, tên này phải là duy nhất cho loại tài nguyên này trong cùng một mô-đun – nhiều hơn nữa trên các mô-đun sau.

- Cuối cùng, chúng ta có một khối chứa một loạt các đối số được sử dụng như một đặc tả tài nguyên. Một điểm chính về tài nguyên là sau khi được tạo, chúng ta có thể sử dụng các biểu thức để truy vấn các thuộc tính của chúng. Ngoài ra, và quan trọng không kém, chúng ta có thể sử dụng các thuộc tính đó làm đối số cho các tài nguyên khác .

- Để minh họa cách hoạt động của điều này, hãy mở rộng ví dụ trước bằng cách tạo phiên bản EC2 của chúng ta trong VPC không mặc định (Virtual Private Cloud):

    ```

    resource "aws_instance" "web" {
      ami = "some-ami-id"
      instance_type = "t2.micro"
      subnet_id = aws_subnet.frontend.id
    }
    resource "aws_subnet" "frontend" {
      vpc_id = aws_vpc.apps.id
      cidr_block = "10.0.1.0/24"
    }
    resource "aws_vpc" "apps" {
      cidr_block = "10.0.0.0/16"
    }

    ```
- Ở đây, chúng ta sử dụng thuộc tính id từ tài nguyên VPC của chúng ta làm giá trị cho đối số vpc_id của client interface . Tiếp theo, tham số id của nó  trở thành đối số cho cá thể EC2. Xin lưu ý rằng cú pháp cụ thể này yêu cầu Terraform phiên bản 0.12 trở lên . Các phiên bản trước sử dụng cú pháp “$ {expression}” rườm rà hơn , cú pháp này vẫn có sẵn nhưng được coi là kế thừa.

- Ví dụ này cũng cho thấy một trong những điểm mạnh của Terraform: bất kể thứ tự mà chúng ta khai báo tài nguyên trong dự án của mình, nó sẽ tìm ra thứ tự chính xác mà nó phải tạo hoặc cập nhật chúng dựa trên biểu đồ phụ thuộc mà nó xây dựng khi phân tích cú pháp chúng.

### count và for_each Đối số Meta

- Các đối số meta count và for_each cho phép chúng ta tạo nhiều phiên bản của bất kỳ tài nguyên nào. Sự khác biệt chính giữa chúng là số đếm mong đợi một số không âm, trong khi for_each  chấp nhận một danh sách hoặc bản đồ các giá trị.

- Ví dụ: hãy sử dụng số đếm để tạo một số phiên bản EC2 trên AWS:

    ```
    resource "aws_instance" "server" {
      count = var.server_count 
      ami = "ami-xxxxxxx"
      instance_type = "t2.micro"
      tags = {
        Name = "WebServer - ${count.index}"
      }
    }

    ```
- Trong một tài nguyên sử dụng số đếm , chúng ta có thể sử dụng  đối tượng đếm trong các biểu thức . Đối tượng này chỉ có một thuộc tính:  chỉ mục , chứa chỉ mục (dựa trên không) của mỗi cá thể.

- Tương tự như vậy, chúng ta có thể sử dụng đối số meta for_each để tạo các trường hợp đó dựa trên map:

    ```
    variable "instances" {
      type = map(string)
    }
    resource "aws_instance" "server" {
      for_each = var.instances 
      ami = each.value
      instance_type = "t2.micro"
      tags = {
        Name = each.key
      }
    }

    ```
- Lần này, chúng ta đã sử dụng map từ nhãn đến tên AMI (Amazon Machine Image) để tạo máy chủ của mình. Trong tài nguyên của mình, chúng ta có thể sử dụng từng đối tượng, đối tượng này cho phép chúng ta truy cập vào khóa và giá trị  hiện tại cho một phiên bản cụ thể.

- Một điểm chính về count  và for_each là, mặc dù chúng ta có thể gán các biểu thức cho chúng, Terraform phải có khả năng phân giải các giá trị của chúng trước khi thực hiện bất kỳ hành động tài nguyên nào . Do đó, chúng ta không thể sử dụng một biểu thức phụ thuộc vào các thuộc tính đầu ra từ các tài nguyên khác.

### Data Sources

- Data Source hoạt động khá nhiều dưới dạng tài nguyên “chỉ đọc” , theo nghĩa là chúng ta có thể lấy thông tin về những nguồn hiện có nhưng không thể tạo hoặc thay đổi chúng. Chúng thường được sử dụng để tìm nạp các tham số cần thiết để tạo các tài nguyên khác.

- Ví dụ điển hình là nguồn dữ liệu aws_ami có sẵn trong nhà cung cấp AWS mà chúng ta sử dụng để khôi phục các thuộc tính từ AMI hiện có:

    ```
    data "aws_ami" "ubuntu" {
      most_recent = true
      filter {
        name   = "name"
        values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
      }
      filter {
        name   = "virtualization-type"
        values = ["hvm"]
      }
      owners = ["099720109477"] # Canonical
    }

    ```
- Ví dụ này định nghĩa một Data Source được gọi là “ubuntu” truy vấn AMI và trả về một số thuộc tính liên quan đến hình ảnh được định vị. Sau đó, chúng ta có thể sử dụng các thuộc tính đó trong các định nghĩa tài nguyên khác, thêm tiền tố data vào tên thuộc tính:
    ```
    resource "aws_instance" "web" {
      ami = data.aws_ami.ubuntu.id 
      instance_type = "t2.micro"
    }

    ```

### State

- State (tệp trạng thái) của một dự án Terraform là một tệp lưu trữ tất cả thông tin chi tiết về tài nguyên đã được tạo trong ngữ cảnh của một dự án nhất định . Ví dụ: nếu chúng ta khai báo tài nguyên azure_resourcegroup trong dự án của mình và chạy Terraform, tệp trạng thái sẽ lưu trữ định danh của nó.

- Mục đích chính của tệp trạng thái là cung cấp thông tin về các tài nguyên hiện có, vì vậy khi chúng ta sửa đổi định nghĩa tài nguyên của mình, Terraform có thể tìm ra những gì nó cần làm.

- Một điểm quan trọng về tệp trạng thái là chúng có thể chứa thông tin nhạy cảm . Ví dụ bao gồm mật khẩu ban đầu được sử dụng để tạo cơ sở dữ liệu, khóa riêng tư, v.v.

- Terraform sử dụng khái niệm backend để lưu trữ và truy xuất các tệp trạng thái. Phần backend mặc định là phần local backend , sử dụng một tệp trong thư mục gốc của dự án làm vị trí lưu trữ. Chúng ta cũng có thể định cấu hình một backend từ xa thay thế bằng cách khai báo nó trong một khối terraform trong một trong các tệp .tf của dự án :

    ```
    terraform {
      backend "s3" {
        bucket = "some-bucket"
        key = "some-storage-key"
        region = "us-east-1"
      }
    }

    ```

### Mô-đun

- Mô-đun Terraform là tính năng chính cho phép chúng ta sử dụng lại các định nghĩa tài nguyên trên nhiều dự án hoặc đơn giản là có một tổ chức tốt hơn trong một dự án duy nhất . Điều này giống như những gì chúng ta làm trong lập trình tiêu chuẩn: thay vì một tệp duy nhất chứa tất cả code, chúng ta tổ chức code của mình trên nhiều package và file.

- Mô-đun chỉ là một thư mục chứa một hoặc nhiều tệp định nghĩa tài nguyên. Trên thực tế, ngay cả khi chúng ta đặt tất cả code của mình trong một tệp / thư mục duy nhất, chúng ta vẫn đang sử dụng các mô-đun – trong trường hợp này, chỉ một. Điểm quan trọng là các thư mục con không được bao gồm như một phần của mô-đun. Thay vào đó, mô-đun mẹ phải bao gồm chúng một cách rõ ràng bằng cách sử dụng khai báo mô-đun :

    ```
    module "networking" {
      source = "./networking"
      create_public_ip = true
    }

    ```
- Ở đây, chúng ta đang tham chiếu đến một mô-đun nằm tại thư mục con “networking” và chuyển một tham số duy nhất cho nó – một giá trị boolean trong trường hợp này.

- Điều quan trọng cần lưu ý là trong phiên bản hiện tại, Terraform không cho phép sử dụng count và for_each để tạo nhiều phiên bản của một mô-đun.

### Biến đầu vào (input variable)

- Bất kỳ mô-đun nào, kể cả mô-đun trên cùng hoặc mô-đun chính, đều có thể xác định một số biến đầu vào bằng cách sử dụng các định nghĩa khối biến  :

    ```
    
    variable "myvar" {
      type = string
      default = "Some Value"
      description = "MyVar description"
    }

    ```
- Một biến có một kiểu ,  có thể là một string , set hoặc map , trong số những biến khác. Nó cũng có thể có giá trị và mô tả mặc định. Đối với các biến được xác định ở mô-đun cấp cao nhất, Terraform sẽ chỉ định các giá trị thực tế cho một biến bằng cách sử dụng một số nguồn:

    - -var tùy chọn dòng lệnh
    
    - .tfvar tệp, sử dụng các tùy chọn dòng lệnh hoặc quét các tệp / vị trí nổi tiếng
    
    - Các biến môi trường bắt đầu bằng TF_VAR_
    
    - Giá trị mặc định của biến , nếu có

- Đối với các biến được xác định trong các mô-đun lồng nhau hoặc bên ngoài, bất kỳ biến nào không có giá trị mặc định phải được cung cấp bằng cách sử dụng các đối số trong  tham chiếu mô-đun . Terraform sẽ tạo ra lỗi nếu chúng ta cố gắng sử dụng mô-đun yêu cầu giá trị cho biến đầu vào nhưng chúng ta không cung cấp được.

- Sau khi được xác định, chúng ta có thể sử dụng các biến trong biểu thức bằng cách sử dụng tiền tố var :

    ```
    resource "xxx_type" "some_name" {
      arg = var.myvar
    }

    ```
### Giá trị đầu ra (output value)

- Theo thiết kế, người tiêu dùng của mô-đun không có quyền truy cập vào bất kỳ tài nguyên nào được tạo trong mô-đun. Tuy nhiên, đôi khi chúng ta cần một số thuộc tính đó để sử dụng làm đầu vào cho mô-đun hoặc tài nguyên khác. Để giải quyết những trường hợp đó, một mô-đun có thể xác định các khối output  hiển thị một tập hợp con của các tài nguyên đã tạo :

    ```
    output "web_addr" {
      value = aws_instance.web.private_ip
      description = "Web server's private IP address"
    }
    ```

- Ở đây chúng ta đang xác định giá trị đầu ra có tên “web_addr” chứa địa chỉ IP của phiên bản EC2 mà mô-đun của chúng ta đã tạo. 
Bây giờ bất kỳ mô-đun nào tham chiếu đến mô-đun của chúng ta đều có thể sử dụng giá trị này trong các biểu thức dưới dạng module.module_name.web_addr ,  trong đó  module_name là tên chúng ta đã sử dụng trong khai báo mô-đun tương ứng  .

### Biến cục bộ

- Các biến cục bộ hoạt động giống như các biến tiêu chuẩn, nhưng phạm vi của chúng bị giới hạn trong mô-đun nơi chúng được khai báo . Việc sử dụng các biến cục bộ có xu hướng giảm sự lặp lại code, đặc biệt là khi xử lý các giá trị đầu ra từ các mô-đun:

    ```
    locals {
      vpc_id = module.network.vpc_id
    }
    module "network" {
      source = "./network"
    }
    module "service1" {
      source = "./service1"
      vpc_id = local.vpc_id
    }
    module "service2" {
      source = "./service2"
      vpc_id = local.vpc_id
    }
    ```
- Ở đây, biến cục bộ vpc_id nhận giá trị của một biến đầu ra từ mô-đun network . Sau đó, chúng ta chuyển giá trị này làm đối số cho cả mô-đun service1 và service2 .

### Workspaces

- Workspaces Terraform cho phép chúng ta giữ nhiều tệp trạng thái (state file) cho cùng một dự án. Khi chúng ta chạy Terraform lần đầu tiên trong một dự án, tệp trạng thái được tạo sẽ đi vào không gian làm việc default . Sau đó, chúng ta có thể tạo workspace mới bằng lệnh terraform workspace new, tùy chọn cung cấp tệp trạng thái hiện có làm tham số.

- Chúng ta có thể sử dụng workspace việc khá nhiều như chúng ta sử dụng các nhánh trong một VCS thông thường . Ví dụ: chúng ta có thể có một workspace cho mỗi môi trường đích – DEV, QA, PROD – và bằng cách chuyển đổi workspace, chúng ta có thể terraform apply khi thêm tài nguyên mới.

- Với cách thức hoạt động, workspace là một lựa chọn tuyệt vời để quản lý nhiều phiên bản nếu bạn muốn – của cùng một nhóm cấu hình. Đây là một tin tuyệt vời cho tất cả những ai đang phải đối phó với vấn đề “works in my environment” khét tiếng, vì nó cho phép chúng ta đảm bảo rằng tất cả các môi trường đều giống nhau.

- Trong một số trường hợp, có thể thuận tiện để tắt tính năng tạo một số tài nguyên dựa trên workspace cụ thể mà chúng ta đang nhắm mục tiêu. Đối với những trường hợp đó, chúng ta có thể sử dụng  biến xác định trước terraform.workspace . Biến này chứa tên của workspace hiện tại và chúng ta có thể sử dụng nó như bất kỳ biến nào khác trong các biểu thức.