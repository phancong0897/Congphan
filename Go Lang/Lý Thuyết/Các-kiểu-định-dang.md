# Các kiểu định dạng trong GO

### Định dạng chung

- %v - In giá trị ở định dạng mặc định.

- %+v - In giá trị ở định dạng mặc định, bao gồm tên trường cho các struct.

- %#v - In giá trị ở định dạng Go-syntax, tương tự như cách mà giá trị có thể được viết trong mã Go.

- %T - In kiểu của giá trị.

- %% - In dấu phần trăm (không lấy giá trị nào).

### Định dạng boolean

- %t - In giá trị boolean (true hoặc false).

### Định dạng số nguyên

- %b - In số nguyên ở dạng nhị phân.

- %c - In ký tự từ mã Unicode tương ứng.

- %d - In số nguyên ở dạng thập phân.

- %o - In số nguyên ở dạng bát phân (octal).

- %O - In số nguyên ở dạng bát phân với tiền tố 0o.

- %q - In ký tự trong dấu nháy đơn với ký tự đặc biệt được thoát.

- %x - In số nguyên ở dạng thập lục phân (hexadecimal), chữ thường.

- %X - In số nguyên ở dạng thập lục phân, chữ hoa.

- %U - In số nguyên ở dạng Unicode (U+1234).

- %d - In số nguyên ở dạng cơ số 10 (decimal).

- %#x - In số nguyên ở dạng thập lục phân với tiền tố 0x.

### Định dạng số thực và số phức

- %b - In số thực ở dạng khoa học với mũ là số nguyên bậc 2.

- %e - In số thực ở dạng khoa học (ví dụ: -1234.456e+78).

- %E - In số thực ở dạng khoa học (ví dụ: -1234.456E+78).

- %f - In số thực ở dạng dấu chấm động (ví dụ: 123.456).

- %F - Tương tự %f.

- %g - In số thực ở dạng gọn gàng hơn giữa dạng dấu chấm động và dạng khoa học.

- %G - Tương tự %g.

### Định dạng chuỗi và byte slice

- %s - In chuỗi hoặc byte slice không qua kiểm tra.

- %q - In chuỗi có dấu nháy kép với ký tự đặc biệt được thoát.

- %x - In chuỗi hoặc byte slice ở dạng thập lục phân, mỗi byte được in thành 2 ký tự.

- %X - Tương tự %x nhưng với chữ hoa.

### Định dạng con trỏ

- %p - In con trỏ ở dạng thập lục phân với tiền tố 0x.

### Định dạng các giá trị không xác định

- %v - In giá trị ở định dạng mặc định.

- %+v - In giá trị ở định dạng mặc định với tên trường cho các struct.

- %#v - In giá trị ở định dạng Go-syntax.

### Ký tự định dạng đặc biệt khác

- %t - In giá trị boolean.

- %T - In kiểu của giá trị.