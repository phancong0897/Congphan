# Terraform variable type

### Terraform Variable

Terraform hỗ trợ 2 loại Biến (Variable) trong quá trình sử dụng Terraform gồm: Input và Output

- Input Variable : được sử dụng như một cách để truyền tham số giá trị cần cấu hình trong hạ tầng của bạn.

- Output Variable : ở chiều ngược lại, thì Output Variable giúp hiển thị thông tin cần thiết về hạ tầng của bạn sau khi triển khai hạ tầng.

### 1. Input Variable là gì ?

- Sử dụng biến là một khái niệm phổ biến trong Terraform. Thay vì phải cấu hình các giá trị thông tin trực tiếp trong các file cấu hình resource Terraform khác nhau, thì bạn có thể quản lý tập trung thông tin, truy vấn thông tin sử dụng từ biến, biến sẽ được truyền vào theo các cách khác nhau để linh động khởi tạo hạ tầng dịch vụ.

- Ví dụ xài giá trị trực tiếp trong resource AWS EC2 Instance Terraform :

    ```

    resource "aws_instance" "web-server" {
      ami           = "ami-0b4dd9d65556cac22"
      instance_type = "t2.medium"
      monitoring    = true

      tags = {
        Name = "ec2.test.cuongquach.com"
        Terraform = "true"
        Environment = "dev"
      }
    }

    ```
- Ví dụ xài giá trị qua biến trong resource AWS EC2 Instance Terraform :

    ```
    + File variables.tf

    variable "set_enabled" {
        default = false
    }

    variable "list-ami-id" {
        type    = "list"
        default = ["ami-0b4dd9d65556cac22", "ami-0c47djd6581jdac1q"]
    }

    variable "ec2_instance_type" {
        type = "string"
        description = "The AWS EC2 Instance type used to deploy resources"
        default = "t2.medium"
    }

    ```
    ```
    + File aws-ec2.tf

    provider "aws" {
      access_key = "xxx"
      secret_key = "xxx"
      region     = "ap-southeast-1"
    }

    resource "aws_instance" "web-server" {
      ami           = "${var.list-ami-id[0]}"
      instance_type = "${var.ec2_instance_type}"
      monitoring    = "${var.set_enabled}"

      tags = {
        Name = "ec2.test.cuongquach.com"
        Terraform = "true"
        Environment = "dev"
      }
    }
    ```

- Vậy Input Variable, được coi như các tham số truyền vào cho một Terraform resource sử dụng. Tham số truyền vào sẽ được khai thác qua các biến đã được định nghĩa sẵn và giá trị truyền vào thông qua Command Line hoặc biến môi trường.

- Tại sao phải định nghĩa Input Variable ?

    - Vì hạ tầng dịch vụ rất là quan trọng ở tính đúng đắn và toàn vẹn dữ liệu. Nên khi bạn muốn sử dụng bất kì Input Variable nào trong resource Terraform thì phải khai báo trước thông tin biến đó gồm kiểu dữ liệu, giá trị mặc định,.. thì mới được sử dụng.
    
    - Chúng ta định nghĩa kiểu dữ liệu Input Variable và dữ liệu mặc định cho Input Variable. Và không được gán giá trị cho kiểu dữ liệu đó khi xài chung cú pháp định nghĩa Input Variable.

- Cấu hình Input Variable ở đâu ?

    - Theo cú pháp khuyến nghị của Terraform thì bạn nên tạo một file có tên là ‘variables.tf‘ ở cùng cấp thư mục cấu hình Terraform để khai báo các biến của bạn, nhằm quản lý tốt hơn.

    - Nhưng điều đó không có nghĩa là bạn buộc phải đặt tên ‘variables.tf’. Bạn hoàn toàn có thể đặt tên file khác, vì đơn giản Terraform sẽ load toàn bộ file có hậu tố *.tf khi chạy. Nên chỉ cần file nào có cú pháp khai báo liên quan đến Biến (variables) thì Terraform sẽ load thông tin về biến đó.

- Cú pháp khai báo Input Variable:

    - Block ‘variable‘ giúp ta cấu hình định nghĩa cấu trúc một biến đầu vào (input variable).
    
    - Trong block { } , chỉ cho phép 3 tham số được sử dụng để cấu hình là : “type” , “default” và “description“.
    
    - Không được sử dụng các tên biến là : source , version , providers . Vì các tên này được Terraform sử dụng cho các chức năng hoạt động nội bộ.

    ```

    variable "NAME_VARIABLE" {
        type = ""
        description = ""
        default = ""
    }
    ```

    - Ví dụ:

    ```
    variable "key" {
      type = "string"
    }

    variable "images" {
      type = "map"

      default = {
        us-east-1 = "image-1234"
        us-west-2 = "image-4567"
      }
    }

    variable "zones" {
      default = ["us-east-1a", "us-east-1b"]
    }
    ```

- Input Variable trong Terraform hỗ trợ nhiều kiểu dữ liệu (type) như sau :

    - string (chuỗi)
    
    - number (số)
    
    - boolean (true/false)
    
    - list (danh sách)
    
    - set
    
    - map
    
    - object
    
    - tuple

    Cũng như hỗ trợ giá trị mặc định cho biến (default value) trong trường hợp không có giá trị nào được đưa vào trong lúc khởi chạy Terraform triển khai hạ tầng.

-  Cú pháp sử dụng Input Variable

    - Với ${<NAME_VARIABLE>} khớp với tên biến đã định nghĩa ở block ‘variable { }‘.

        ` ${var.<NAME_VARIABLE>} `
    
    - Nếu xài biến trong ‘expression‘ của Terraform thì như dưới :

        ` var.<NAME_VARIABLE> `

### 2. Các kiểu dữ liệu Input Variable trong Terraform

- Chúng ta sẽ tìm hiểu các loại dữ liệu phổ biến thường được khai báo sử dụng trong Input Variale Terraform gồm:

    - string
    
    - number
    
    - boolean
    
    - list
    
    - set
    
    - map


### 2.1 String

- Kiểu này dữ liệu ‘string‘ giúp ta lưu trữ giá trị thông tin biến ở dạng chuỗi.

    ```
    variable "ec2_instance_type" {
        type = "string"
        description = "The AWS EC2 Instance type used to deploy resources"
        default = "t2.micro"
    }
    ```

- Trong cấu hình resource bạn sẽ truy vấn thông tin biến ‘ec_instance_type‘ như sau.

    ```
    resource "aws_instance" "web-server" {
      ami           = "ami-0b4dd9d65556cac22"
      instance_type = "${var.ec2_instance_type}"

      tags = {
        Name = "ec2.test.cuongquach.com"
        Terraform = "true"
        Environment = "dev"
      }
    }

    ```

- Nếu bạn muốn cấu hình biến với dữ liệu nhiều dòng, có thể làm như dưới :

    ```

    variable "long_key" {
      type = "string"
      default = <<EOF
    This is a long key.
    Running over several lines.
    EOF
    }

    ```

### 2.2 Number

- Kiểu dữ liệu ‘number‘ giúp ta khai báo lưu trữ giá trị thông tin ở dạng số. Ví dụ : 15 hoặc 6.555 . Khi khai báo kiểu ‘number’ thì không cần dấu ngoặc kép “…” .

    ```
    variable "ec2_cpu_core_count" {
        type = number
        default = 2
    }

    ```
- Truy vấn thông tin kiểu dữ liệu Number như dưới. Ví dụ: <cpu_core_count>

    ```
    resource "aws_instance" "web-server" {
      instance_type = "${var.ec2_instance_type}"
      cpu_core_count = "${var.ec2_cpu_core_count}"

      tags = {
        Name = "ec2.test.cuongquach.com"
        Terraform = "true"
        Environment = "dev"
      }
    }

    ```

### 2.3 Boolean

- Boolean thì ai cũng biết đó là đại diện cho “True” và “False“. Thường được dùng cho việc xử lý điều kiện logic. Khi khai báo kiểu ‘boolean‘ thì không cần chỉ định loại cho nó.
    ```

    variable "set_monitoring_enabled" {
        default = false
    }

    ```
- Ta truy vấn thông tin kiểu dữ liệu Boolean như dưới. Ví dụ: <monitoring>

    ```
    resource "aws_instance" "web-server" {
      ami           = "${var.list-ami-id[0]}"
      instance_type = "${var.ec2_instance_type}"
      monitoring    = "${set_monitoring_enabled}"

      tags = {
        Name = "ec2.test.cuongquach.com"
        Terraform = "true"
        Environment = "dev"
      }
    }
    ```

### 2.4 List

- List có thể xem như một danh sách các phần tử , một dạng array vậy. Mỗi phần tử trong danh sách này sẽ được truy xuất dựa vào vị trí index của List.
    
    ```
    variable "list-ami-id" {
        type    = "list"
        default = ["ami-0b4dd9d65556cac22", "ami-0c47djd6581jdac1q"]
    }

    ```
- Ta sẽ truy vấn thông tin của kiểu dữ liệu List theo vị trí index như sau . Ví dụ: <ami>

    ```
    resource "aws_instance" "web-server" {
      ami           = "${var.list-ami-id[0]}"
      instance_type = "${var.ec2_instance_type}"

      tags = {
        Name = "ec2.test.cuongquach.com"
        Terraform = "true"
        Environment = "dev"
      }
    }

    ```
### 2.5 Map

- Map là kiểu dữ liệu biến tập hợp thông tin theo dạng key/value. Loại dữ liệu này khá phù hợp để bạn có thể lựa chọn thông tin giá trị dưới theo các thông tin đã được định nghĩa sẵn. Ví dụ: chúng ta phân loại cấu hình server dựa theo mức giá :

    ```
    variable "plans" {
        type = "map"
        default = {
            "5USD"  = "1xCPU-1GB"
            "10USD" = "1xCPU-2GB"
            "20USD" = "2xCPU-4GB"
        }
    }
    ```

- Bạn có thể truy vấn giá trị cấu hình plan mà chúng ta đã định nghĩa trong kiểu dữ liệu map.

    `plan = "${var.plans["5USD"]}"`

- Giá trị trên cũng có thể dùng tiếp tục để truy vấn thông tin ở một biến Map khác. Ví dụ: với mức cấu hình phần cứng giá 5USD bạn sẽ quyết định chọn dung lượng ổ cứng tương đương mà bạn đã định nghĩa sẵn.

    ```
    variable "storage_sizes" {
        type = "map"
        default = {
            "1xCPU-1GB"  = "25"
            "1xCPU-2GB"  = "50"
            "2xCPU-4GB"  = "80"
        }
    }

    ```
- Với cú pháp dưới bạn sẽ truyền biến ở một dữ liệu biến Map ‘plan‘ sang tìm kiếm ở một dữ liệu biến Map ‘storage_sizes‘.

    `size = lookup(var.storage_sizes, var.plans["5USD"]) `

### 2.6 Set

- Set là tập hợp các giá trị phần tử như danh sách (list), nhưng có thêm thuộc tính là “duy nhất”. Tức đảm bảo trong danh sách phần tử sẽ không có 2 phần tử có giá trị giống nhau.

- Ví dụ logic:

    list_a = ["cuongqc", "cuongqc", "cuongqc", "demo2", "demo3"]
    
    thì khi gán kiểu dữ liệu Set vào thì sẽ thành.

    set_a = ["cuongqc", "demo2", "demo3"]


### 3. Sử dụng Input Variable trong Terraform như thế nào ?

- Ở phần trên chúng ta đã học cách làm chủ việc khai báo định nghĩa các biến sẽ sử dụng trong Terraform resource. Ở phần này chúng ta sẽ tìm hiểu cách truyền giá trị biến mà chúng ta mong muốn vào quá trình chạy Terraform resource.

### 3.1 Sử dụng default Input Variable

- Như bạn đã tìm hiểu ở trên, chúng ta có cấu hình giá trị “default” tức là giá trị mặc định của biến đó sẽ được sử dụng, nếu không được truyền thông tin giá trị khác vào trong quá trình chạy Terraform khởi tạo hạ tầng.
    ```
    variable "public_key" {
        type = "string"
        default = "ssh-rsa terraform_public_key"
    }
    ```
### 3.2 Truyền biến từ Command Line

- Bạn có thể truyền giá trị biến trực tiếp trong lúc chạy command line ‘terraform‘ lúc chạy ‘plan‘ và ‘apply‘. Với option ‘-var‘ sẽ giúp bạn khai báo biến mong muốn.

    - terraform plan -var='set_enabled=true' -var='ec2_instance_type=t2.micro' -var='list-ami-id=["ami-0b4dd9d65556cac22", "ami-0c47djd6581jdac1q"]'

### 3.3 Truyền biến từ file định nghĩa biến riêng

- Thay vì phải truyền biến từ các option qua lệnh Command Line, thì bạn có thể chuẩn bị sẵn một file text chứa đầy đủ thông tin biến cần truyền vào gồm tên biến và giá trị biến.

- Cú pháp sử dụng:

- File định nghĩa biến và giá trị phải có hậu tố tên file là : .tfvars hoặc .tfvars.json

- Sử dụng option ‘-var-file’ để load file được chỉ định.

    ` terraform apply -var-file="dev-cuongquach.tfvars" `

- Nội dung file *.tfvars cũng đơn giản thôi, tên biến và giá trị biến.

    ```
    image_id = "ami-abc123"
    availability_zone_names = [
      "us-east-1a",
      "us-west-1c",
    ]
    ```

- Terraform cũng sẽ tự động tìm kiếm các file có cú pháp tên cụ thể để load thông tin biến và giá trị của biến, nếu bạn lười chỉ định option ‘-var-file‘.

    - Tên file phải là : terraform.tfvars hoặc terraform.tfvars.json

    - Hoặc bất kì file gì có tên kết thúc bằng : .auto.tfvars hoặc .auto.tfvars.json

### 3.4 Biến môi trường

- Đơn giản nhanh gọn lẹ với việc bạn tạo biến môi trường khi chạy chương trình ‘terraform‘. Lúc này ‘terraform‘ tự kiểm tra các biến môi trường mà Terraform thấy được và sử dụng chúng.

    ```
    # export TF_VAR_image_id=ami-abc123
    # terraform plan
    ```

### 3.5 Thứ tự load biến khi chạy Terraform

- Thứ tự load biến khi chạy Terraform sẽ theo trình tự như sau :

    - Biến môi trường
    
    - File terraform.tfvars , nếu có tồn tại
    
    - File terraform.tfvars.json , nếu có tồn tại
    
    - Bất kì file *.auto.tfvars hoặc *.auto.tfvars.json

Các biến được chỉ định từ option ‘-var’ hoặc ‘-var-file’, theo thứ tự chỉ định khi chạy chương trình lệnh.

### 4. Output Variable

- Output variable cung cấp cho chúng ta một cách thức để biết được các thông tin hữu ích về hạ tầng của bạn sau khi triển khai bằng Terraform. Nếu bạn để ý, thì rất nhiều thông tin chi tiết về dịch vụ hoặc máy chủ cloud chỉ có thể biết được sau khi Terraform đã khởi tạo hoàn tất. Sử dụng biến Output giúp chúng ta trích xuất các giá trị thông tin về dịch vụ hoặc máy chủ mà chúng ta đã khởi tạo hoàn tất và hiển thị ra màn hình.

- Khai báo Output variable vô cùng đơn giản. Bạn có thể tạo một file tên ‘output.tf‘ để chứa các khai báo biến Output. Kế đến những gì bạn cần làm là định nghĩa 1 cái tên cho output và giá trị mà output đó cần thể hiện. Ví dụ, bạn cần Terraform hiển thị thông tin địa chỉ IP của server sau khi triển khai xong.

    ```
    output "instance_ip_addr" {
      value = aws_instance.server.private_ip
    }
    ```

- Sau khi Terraform chạy xong quá trình ‘apply‘, thì sẽ hiển thị thông tin output mà bạn định nghĩa ở cuối kết quả dòng lệnh. Giờ ta làm ví dụ đơn giản, là hiển thị thông tin ‘ip_private‘ và ‘ip_public‘ của máy chủ EC2 thử nhé.

    ```
    # vi variables.tf
    output "public_ip" {
        value = aws_instance.server.public_ip
    }

    output "private_ip" {
        value = aws_instance.server.private_ip
    }
    # vi ec2.tf
    provider "aws" {
      access_key = "xxx"
      secret_key = "xxx"
      region     = "ap-southeast-1"
    }

    resource "aws_instance" "web-server" {
      ami           = "${var.list-ami-id[0]}"
      instance_type = "${var.ec2_instance_type}"

      tags = {
        Name = "ec2.test.cuongquach.com"
        Terraform = "true"
        Environment = "dev"
      }
    }
    ```
- Giờ ta chạy lệnh ‘terraform apply‘ để khởi tạo một máy chủ EC2 cơ bản. Và bạn sẽ thấy output ở cuối dòng kết quả thực hiện.

    ```
    # terraform apply
    aws_instance.web-server: Creating...
    aws_instance.web-server: Still creating... [10s elapsed]
    aws_instance.web-server: Still creating... [20s elapsed]
    aws_instance.web-server: Still creating... [30s elapsed]
    aws_instance.web-server: Creation complete after 31s [id=i-078b2e78ae7ecf09e]
    
    Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
    
    Outputs:
    
    private_ip_address = 172.31.5.153
    public_ip_address = 3.0.100.169

    ```