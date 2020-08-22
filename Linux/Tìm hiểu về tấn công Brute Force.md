# Tìm hiểu về tấn công Brute Force

# Mục lục

## [1. Tấn công Brute Force là gì?](https://github.com/phancong0897/Congphan/blob/master/T%C3%ACm%20hi%E1%BB%83u%20v%E1%BB%81%20t%E1%BA%A5n%20c%C3%B4ng%20Brute%20Force.md#1-t%E1%BA%A5n-c%C3%B4ng-brute-force-l%C3%A0-g%C3%AC-1)

## [2. Tốc độ của máy tính và vấn để mật khẩu trong cuộc tấn công Brute Force](https://github.com/phancong0897/Congphan/blob/master/T%C3%ACm%20hi%E1%BB%83u%20v%E1%BB%81%20t%E1%BA%A5n%20c%C3%B4ng%20Brute%20Force.md#2-t%E1%BB%91c-%C4%91%E1%BB%99-c%E1%BB%A7a-m%C3%A1y-t%C3%ADnh-v%C3%A0-v%E1%BA%A5n-%C4%91%E1%BB%83-m%E1%BA%ADt-kh%E1%BA%A9u-trong-cu%E1%BB%99c-t%E1%BA%A5n-c%C3%B4ng-brute-force-1)

## [3. Cách phòng chống và bảo vệ để tránh khỏi các cuộc tấn công Brute Force](https://github.com/phancong0897/Congphan/blob/master/T%C3%ACm%20hi%E1%BB%83u%20v%E1%BB%81%20t%E1%BA%A5n%20c%C3%B4ng%20Brute%20Force.md#3-c%C3%A1ch-ph%C3%B2ng-ch%E1%BB%91ng-v%C3%A0-b%E1%BA%A3o-v%E1%BB%87-%C4%91%E1%BB%83-tr%C3%A1nh-kh%E1%BB%8Fi-c%C3%A1c-cu%E1%BB%99c-t%E1%BA%A5n-c%C3%B4ng-brute-force-1)

## [4. Sử dụng tools medusa tấn công brute force](https://github.com/phancong0897/Congphan/blob/master/T%C3%ACm%20hi%E1%BB%83u%20v%E1%BB%81%20t%E1%BA%A5n%20c%C3%B4ng%20Brute%20Force.md#4-s%E1%BB%AD-d%E1%BB%A5ng-tools-medusa-t%E1%BA%A5n-c%C3%B4ng-brute-force-1)

## [Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/T%C3%ACm%20hi%E1%BB%83u%20v%E1%BB%81%20t%E1%BA%A5n%20c%C3%B4ng%20Brute%20Force.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o-1)


## 1. Tấn công Brute Force là gì?

- Tấn công Brute Force là một loại tấn công mạng, trong đó bạn có một phần mềm, xoay vòng các ký tự khác nhau, kết hợp để tạo ra một mật khẩu đúng. Phần mềm Brute Force Attack password cracker đơn giản sẽ sử dụng tất cả các kết hợp có thể để tìm ra mật khẩu cho máy tính hoặc máy chủ mạng. Nó rất đơn giản và không sử dụng bất kỳ kỹ thuật thông minh nào. Vì phương pháp này chủ yếu dựa trên toán học, phải mất ít thời gian hơn để crack mật khẩu, bằng cách sử dụng các ứng dụng brute force thay vì tìm ra chúng theo cách thủ công. Nói phương pháp này dựa trên toán học vì máy tính làm rất tốt các phép toán và thực hiện chúng trong vài giây, nhanh hơn rất nhiều lần so với bộ não con người (mất nhiều thời gian hơn để tạo ra các sự kết hợp).

- Tấn công Brute Force là tốt hay xấu tùy thuộc vào người sử dụng nó. Nó có thể được bọn tội phạm mạng cố gắng sử dụng để hack vào một máy chủ mạng, hoặc nó có thể được một quản trị viên mạng dùng để xem mạng của mình được bảo mật có tốt không. Một số người dùng máy tính cũng sử dụng các ứng dụng brute force để khôi phục mật khẩu đã quên.

## 2. Tốc độ của máy tính và vấn để mật khẩu trong cuộc tấn công Brute Force

- Nếu mật khẩu của bạn đang sử dụng tất cả các chữ cái thường và không có ký tự đặc biệt hoặc chữ số, chỉ mất 2-10 phút là một cuộc tấn công brute force có thể crack mật khẩu này. Ngược lại, một mật khẩu có sự kết hợp của cả chữ hoa và chữ thường cùng với một vài chữ số (giả sử có 8 chữ số) sẽ mất hơn 14-15 năm để bị crack.

- Nó cũng phụ thuộc vào tốc độ của bộ vi xử lý máy tính, như bao lâu để crack mật khẩu của mạng hoặc đăng nhập bình thường vào một máy tính Windows độc lập.

- Vì vậy, một mật khẩu mạnh mang lại rất nhiều ý nghĩa. Để tạo một mật khẩu thực sự mạnh, bạn có thể sử dụng các ký tự ASCII để tạo mật khẩu mạnh hơn. Ký tự ASCII tham chiếu đến tất cả các ký tự có sẵn trên bàn phím và hơn thế nữa (bạn có thể xem chúng bằng cách nhấn ALT + số (từ 0 đến 255) trên Numpad). Có khoảng 255 ký tự ASCII và mỗi ký tự có một code được đọc bằng máy và chuyển thành nhị phân (0 hoặc 1), sao cho nó có thể được sử dụng bằng máy tính. Ví dụ, code ASCII cho dấu cách là 32. Khi bạn nhập một dấu cách, máy tính đọc nó là 32 và chuyển đổi nó thành nhị phân - sẽ là 10000. Các ký tự 1, 0, 0, 0, 0, 0 được lưu trữ dưới dạng ON, OFF, OFF, OFF, OFF, OFF trong bộ nhớ của máy tính. Điều này không liên quan gì đến brute force, trừ trường hợp bạn sử dụng tất cả các ký tự ASCII. Nếu bạn sử dụng các ký tự đặc biệt trong mật khẩu, tổng thời gian cần để crack mật khẩu có thể lên đến 100 năm.

- Brute Force Password Calculator (link tham khảo: https://www.grc.com/haystack.htm) là nơi bạn có thể kiểm tra xem sẽ mất bao lâu để crack một mật khẩu. Có các tùy chọn khác nhau bao gồm chữ thường, chữ hoa, chữ số và tất cả các ký tự ASCII. Dựa trên những gì bạn đã sử dụng trong mật khẩu, hãy chọn các tùy chọn và nhấp vào nút Calculate để xem cuộc tấn công Brute force sẽ khó khăn như thế nào để crack mật khẩu máy tính hoặc máy chủ của bạn.

## 3. Cách phòng chống và bảo vệ để tránh khỏi các cuộc tấn công Brute Force

- Vì không có logic đặc biệt nào được áp dụng trong các cuộc tấn công Brute Force, ngoại trừ việc thử các kết hợp khác nhau của các ký tự được sử dụng để tạo mật khẩu, nên biện pháp phòng ngừa ở mức rất cơ bản và tương đối dễ dàng.

- Ngoài việc sử dụng phần mềm bảo mật và hệ điều hành Windows được cập nhật đầy đủ, bạn nên sử dụng một mật khẩu mạnh có một số đặc điểm sau:

    - Có ít nhất một chữ hoa

    - Có ít nhất một chữ số

    - Có ít nhất một ký tự đặc biệt

    - Mật khẩu phải có tối thiểu 8-10 ký tự

    - Bao gồm ký tự ASCII, nếu bạn muốn.

- Mật khẩu càng dài thì càng mất nhiều thời gian để crack nó. Nếu mật khẩu của bạn giống như 'PA$$w0rd', sẽ mất hơn 100 năm để crack nó bằng các ứng dụng tấn công brute force hiện có. Xin vui lòng không sử dụng mật khẩu được đề xuất trong ví dụ, vì nó rất dễ dàng bị phá vỡ, bằng cách sử dụng một số phần mềm thông minh, có thể tổng hợp các mật khẩu đề xuất trong các bài viết liên quan đến các cuộc tấn công brute force.

## 4. Sử dụng tools medusa tấn công brute force

### Chuẩn bị:

- Chuẩn bị 2 máy centos7 tạm gọi là Centos1, Centos2

    - Centos 1 có cài đặt fail2ban.

    - Centos2 có cài đặt tool medusa.

- Về máy Centos1 cài đặt fail2ban, ta đã cài đặt ở phần trước, các bạn có thể tham khảo tại 

https://github.com/phancong0897/Congphan/blob/master/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md

- Centos2 cài đặt medusa

    - Cài đặt một số gói cần thiết

        - ` yum update -y `

        - ` yum install epel-release `
    
    - Cài đặt medusa

        - ` yum install medusa `

    - Bạn có thể sử dụng lệnh

        - ` medusa -h `

    - Để xem các option của nó. Dưới đây là một số option mà tôi hay sử dungj:

        - -h : Địa chỉ của máy mục tiêu ta muốn thực hiện brute force ( IP hoặc domain )

        - -H : file chứa danh sách các máy mục tiêu

        - -u : chỉ ra username dùng để kiểm tra
        
        - -U : chỉ ra file chứa danh sách các username
        
        - -p : chỉ ra password dùng để kiểm tra
        
        - -P : chỉ ra file chứa danh sách password dùng để kiểm tra
        
        - -O : chỉ ra file để nó ghi log vào
        
        - -d : hiển thị toàn module mà medusa hỗ trợ
        
        - -M : chỉ ra module bạn muốn brute force ( dùng option -d bên trên để xem toàn bộ module được hỗ trợ nhưng bỏ qua phần đuôi .mod)
        
        - -m : truyền vào các tham số yêu cầu của các module. Option này có thể được sử dụng nhiều lần ( Vd: -m param1 -m param2)
        
        - -n : chỉ ra địa chỉ port
        
        - -s : cho phép sử dụng SSL
        
        - -f : nếu bạn sử đang tấn công cùng lúc cả username và password bạn sử dụng tùy chọn này thì câu lệnh sẽ dừng khi nó tìm thấy thông tin đăng nhập đầu tiên.
        
        - -F : nếu bạn đang thực hiện tấn công đồng thời nhiều máy option này sẽ dừng câu lệnh khi có 1 username và password của bất kì host nào được tìm thấy.
        
        - -q : chỉ ra các tham số sử dụng cho module được chỉ ra
        
        - -Z : nếu bạn đang thực hiện brute force với một dang sách mật khẩu và đã đến mật khảu thứ 100 là 100100 và vì một lý do nào đó bạn nhấn Ctrl C và thoát ra khỏi câu lệnh. Bây giờ bạn muốn scan tiếp từ mật khẩu này bạn sử dụng option -Z và chỉ ra mật khẩu lúc trước đã đến -Z 100100

### Tấn công brute force

- Ở máy Centos 2: ta dùng lệnh để tấn công brute force

    - medusa -h 172.16.2.230 -u root -p passwd.txt -M ssh

<img src="https://imgur.com/eoQtnLA.png">

- Ở Centos1 dùng các lệnh sau để kiểm tra :

    - Để kiểm tra

        - ` fail2ban-client status sshd `

    <img src="https://imgur.com/xfeQDz5.png">

    - Để Để mở khoá một IP các bạn có thể sử dụng lệnh như sau

        - ` fail2ban-client set sshd unbanip IP-Bị-Block `

## Nguồn tham khảo

https://quantrimang.com/cuoc-tan-cong-brute-force-la-gi-157987

https://news.cloud365.vn/brute-force-password-su-dung-medusa/