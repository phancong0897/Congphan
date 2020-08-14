# Tài liệu cài đặt LVM

# Mục lục

[ Cấu hình ]()

[Tạo các partion]()

[ Tạo Physical Volume]()

[ Tạo Volume Group]()

[ Tạo Logical Volume]()

[ Mount và sử dụng]()

[Nguồn tham khảo]()

##  Cấu hình

- Tạo máy ảo Centos 7 trên Vmware.

- Add 3 disk ổ cứng vào máy ảo.

- Sau khi add xong khởi động lại máy ảo , kiểm tra xem máy đã nhận 3 ổ ảo chưa bằng lệnh

    - lsblk

## Tạo các partion

– Tạo các partition cho các ổ mới , bắt đầu từ sdb với lệnh :

    - fdisk /dev/sdb

<img src="https://imgur.com/undefined.png">

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

<img src="https://imgur.com/undefined.png">

## Tạo Physical Volume :

- Từ các partition /dev/sdb1 /dev/sdc1 /dev/sdd1 ta tạo các Physical Volume bằng lệnh sau :
    - pvcreate /dev/sdb1

    - pvcreate /dev/sdc1

    - pvcreate /dev/sdd1

- Kiểm tra bằng lệnh pvs hoặc pvdisplay xem các physical volume đã được tạo chưa :

<img src="https://imgur.com/undefined.png">

## Tạo Volum groupe :
- Sau khi tạo các Physical Volume ta gộp các PV đó thành 1 Volume Group bằng lệnh sau :

    - vgcreate vg-demo1 /dev/sdb1 /dev/sdc1 /dev/sdd1

- Dùng các lệnh vgs hoặc vgdisplay để kiểm tra :

<img src="https://imgur.com/undefined,png">

## Tạo Logical Volume:

- Từ một Volume group , ta tạo các Logical Volume để sử dụng bằng lệnh sau :

    - lvcreate -L 1G -n lv-demo1 vg-demo1

- Trong đó :

    - L : Chỉ ra dung lượng của logical volume

    - n : Chỉ ra tên của logical volume

    - Kiểm tra bằng lệnh lvs hoặc lvdisplay

<img src="https://imgur.com/undefined.png">

- Định dạng Logical Volume:

    - Format các Logical Volume thành các định dạng ext2,ext3,ext4,… :

        - mkfs -t ext4 /dev/vg-demo1/lv-demo1

## Mount và sử dụng:

- Ta tạo một thư mục để mount Logical Volume :

    - mkdir demo1

- Tiến hành mount logical volume lv-demo1 vào thư mục demo1

    - mount /dev/vg-demo1/lv-demo1 demo

- Kiểm tra bằng lệnh df -h

Nguồn tham khảo

https://news.cloud365.vn/lvm-gioi-thieu-ve-logical-volume-manager/