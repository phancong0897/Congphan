# Tài liệu cài đặt LVM

# Mục lục

[ Cấu hình ](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#c%E1%BA%A5u-h%C3%ACnh)

[Tạo các partion](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-c%C3%A1c-partion)

[ Tạo Physical Volume](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-physical-volume-)

[ Tạo Volume Group](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-volum-groupe-)

[ Tạo Logical Volume](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-logical-volume)

[ Mount và sử dụng](phancong0897)

[Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

##  Cấu hình

- Tạo máy ảo Centos 7 trên Vmware.

- Add 3 disk ổ cứng vào máy ảo.

- Sau khi add xong khởi động lại máy ảo , kiểm tra xem máy đã nhận 3 ổ ảo chưa bằng lệnh

    - lsblk

## Tạo các partion

– Tạo các partition cho các ổ mới , bắt đầu từ sdb với lệnh :

    - fdisk /dev/sdb

<img src="https://imgur.com/Ur6Hy80.png">

- Trong đó :

    - Command n để bắt đầu tạo partition.
    - Chọn p để tạo partition primary
    - Chọn 1 để đặt partition number
    - First sector để mặc định
    - Last sector để +1G (có các tùy chọn KB, MB ,GB)
    - Sau khi hoàn thành các bước trên nhấn t để đổi định dạng partition
    - Chọn 8e để thay đổi định dạng partition
    - Sau khi hoàn thành ấn w để lưu và thoát
– Tương tự tạo thêm các partition primary từ sdb ,sdc,sdd.

<img src="https://imgur.com/VspLewq.png">

## Tạo Physical Volume :

- Từ các partition /dev/sdb1 /dev/sdc1 /dev/sdd1 ta tạo các Physical Volume bằng lệnh sau :
    - pvcreate /dev/sdb1

    - pvcreate /dev/sdc1

    - pvcreate /dev/sdd1

- Kiểm tra bằng lệnh pvs hoặc pvdisplay xem các physical volume đã được tạo chưa :

<img src="https://imgur.com/fwPazcP.png">

## Tạo Volum groupe :
- Sau khi tạo các Physical Volume ta gộp các PV đó thành 1 Volume Group bằng lệnh sau :

    - vgcreate vg-demo1 /dev/sdb1 /dev/sdc1 /dev/sdd1

- Dùng các lệnh vgs hoặc vgdisplay để kiểm tra :

<img src="https://imgur.com/ua783j5.png">

## Tạo Logical Volume:

- Từ một Volume group , ta tạo các Logical Volume để sử dụng bằng lệnh sau :

    - lvcreate -L 1G -n lv-demo1 vg-demo1

- Trong đó :

    - L : Chỉ ra dung lượng của logical volume

    - n : Chỉ ra tên của logical volume

    - Kiểm tra bằng lệnh lvs hoặc lvdisplay

<img src="https://imgur.com/LsaQrrs.png">

- Định dạng Logical Volume:

    - Format các Logical Volume thành các định dạng ext2,ext3,ext4,… :

        - mkfs -t ext4 /dev/vg-demo1/lv-demo1

## Mount và sử dụng:

- Ta tạo một thư mục để mount Logical Volume :

    - mkdir demo1

- Tiến hành mount logical volume lv-demo1 vào thư mục demo1

    - mount /dev/vg-demo1/lv-demo1 demo

- Kiểm tra bằng lệnh df -h

- Các phân vùng gắn kết của chúng ta là tạm thời. Nếu hệ điều hành được khởi động lại, các thư mục được gắn kết này sẽ bị mất. Vì vậy, chúng ta cần phải gắn kết vĩnh viễn. Để thực hiện gắn kết vĩnh viễn phải nhập trong tệp /etc/fstab.

    - Để lấy UUID của các phân vùng chúng ta thực hiện như sau:
    
         blkid  /dev/vg-demo1/lv-demo1

<img src="https://imgur.com/09kFerR.png">

    - Sau đó sử dụng lênh vi /etc/fstab để gắn kết

<img src="https://imgur.com/ThWVJBQ.png">

## Nguồn tham khảo

https://news.cloud365.vn/lvm-gioi-thieu-ve-logical-volume-manager/

https://blogd.net/linux/quan-ly-phan-vung-dia-cung-tren-linux/