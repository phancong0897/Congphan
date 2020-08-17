# Tài liệu cài đặt LVM

# Mục lục

[ Cấu hình ](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#c%E1%BA%A5u-h%C3%ACnh)

[Tạo các partion](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-c%C3%A1c-partion)

[ Tạo Physical Volume](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-physical-volume-)

[ Tạo Volume Group](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-volum-groupe-)

[ Tạo Logical Volume](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#t%E1%BA%A1o-logical-volume)

[ Mount và sử dụng](phancong0897)

[Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/C%C3%A0i%20%C4%91%E1%BA%B7t%20LVM.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Cấu hình

- Tạo máy ảo Centos 7 trên Vmware.

- Add 3 disk ổ cứng vào máy ảo.

- Sau khi add xong khởi động lại máy ảo , kiểm tra xem máy đã nhận 3 ổ ảo chưa bằng lệnh

    - lsblk

## 2. Cài đặt

### 2.1 Tạo các partion

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

### 2.2 Tạo Physical Volume :

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

### 2.3 Tạo Logical Volume:

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

### 2.4 Mở rộng Volume Group và thay đổi kích thước các Logical Volume

- Trong ví dụ dưới đây chúng ta sẽ thêm một physical volume có tên /dev/sdd với kích thước 10GB vào volume group vg0, sau đó chúng ta sẽ tăng kích thước của logical volume /projects lên 10GB thực hiện như sau:

- Chạy các lệnh sau để tạo điểm gắn kết:

  - mkdir /projects

  - mkdir /backups

- Chạy lệnh sau để mount:
 
  - mount /dev/vg0/projects /projects/

  - mount /dev/vg0/backups /backups/

- Chạy lệnh sau kiểm tra sử dụng không gian đĩa hệ thống tập tin:

<img src="https://imgur.com/7INSWsk.png">

- Sử dụng lệnh sau để thêm /dev/sdd1 vào volume group vg0:

  - vgextend vg0 /dev/sdd1

<img src="https://imgur.com/WjT2O0e.png">

- Qua lệnh kiểm tra trên chúng ta thấy dung lượng thêm vào của chúng ta là 10GB chúng ta có thể tăng kích thước của logical volume /projects lên 10GB thực hiện như sau:

  - lvextend -l +2000 /dev/vg0/projects

- Sau khi chạy lệnh trên chúng ta cần thay đổi kích thước hệ thống tệp, vì thế chúng ta phải chạy lệnh sau để resize:

  - Đối với file system (ext2, ext3, ext 4): resize2fs

<img src="https://imgur.com/Hm9xgZM.png">

## 2.5 Giảm kích cỡ Logical Volume

- Khi chúng ta muốn giảm dung lượng các Logical Volume. Chúng ta cần phải chú ý vì nó rất quan trọng và có thể bị lỗi trong khi chúng ta giảm dung lượng Logical Volume. Để đảm bảo an toàn khi giảm Logical Volume cần thực hiện các bước sau:

  - Trước khi bắt đầu, cần sao lưu dữ liệu vì vậy sẽ không phải đau đầu nếu có sự cố xảy ra.

  - Để giảm dung lượng Logical Volume, cần thực hiện đầy đủ và cẩn thận 5 bước cần thiết.

  - Khi giảm dung lượng Logical Volume chúng ta phải ngắt kết nối hệ thống tệp trước khi 
  giảm.

- Cần thực hiện cẩn thận 5 bước dưới đây:

  - Ngắt kết nối file system.

  - Kiểm tra file system sau khi ngắt kết nối.

  - Giảm file system.

  - Giảm kích thước Logical Volume hơn kích thước hiện tại.

  - Kiểm tra lỗi cho file system.

  - Mount lại file system và kiểm tra kích thước của nó.

- Ví dụ: Giảm Logical Volume có tên project với kích thước từ 17.81GB giảm xuống còn 10GB mà không làm mất dữ liệu. Chúng ta thực hiện các bước như sau:

  - Bước 1: Umount file system bằng câu lệnh:

    umount /projects/
  
  - Bước 2:  Kiểm tra lỗi file system bằng lệnh e2fsck:

<img src="https://imgur.com/WxRj5uy.png">

  - Bước 3: Giảm kích thước của projects theo kích thước mong muốn.

Giảm Logical Volume có tên projects với kích thước từ 17.81GB giảm xuống còn 10GB chúng ta thực hiện như sau:

Kiểm tra dung lượng của Logical Volume và kiểm tra file system trước khi thực hiện giảm:

 resize2fs /dev/vg0/projects 10G

 - Bước 4: Bây giờ giảm kích thước bằng lệnh lvreduce.

<img src="https://imgur.com/gMMitcI.png">

 - Bước 5: Để đảm bảo an toàn, bây giờ kiểm tra lỗi file system đã giảm

<img src="https://imgur.com/undefined.png">

 - Bước 6: Gắn kết file system và kiểm tra kích thước của nó.

<img src="https://imgur.com/JOiGSE0.png">

### 2.6 Mount và sử dụng:

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

## 3. Snapshots và restore của logical volume

### 3.1 Snapshots

### Tạo Snapshot
- Kiểm tra không gian trống trong volume groupg để tạo snapshot mới chúng ta thực hiện như sau:

<img src="https://imgur.com/m5zalJl.png">

- Qua lệnh kiểm tra trên chúng ta thấy hiện có 2.18GB dung lượng trống còn lại. Vì vậy, tạo một snapshot có tên là snapshot_1 với dung lượng 1GB bằng các lệnh sau:

<img src="https://imgur.com/BglOYkc.png">

- Trong đó ý nghĩa các tuỳ chọn là:

  - s: Tạo snapshot

  - n: Tên cho snapshot

  - Nếu bạn muốn xóa snapshot, có thể dử dụng lệnh lvremove: lvremove /dev/vg0/snapshot_2

- Kiểm tra lại snapshot đã được tạo, chúng ta thực hiện như bên dưới:

<img src="https://imgur.com/IkT7eO5.png">

- Để biết thêm thông tin chi tiết chúng ta sử dụng lệnh sau:

<img src="https://imgur.com/LEcoekg.png">

- Trong đó ý nghĩa các trường của lệnh lvdisplay như sau:

  - LV Name: Tên của Snapshot Logical Volume.

  - VG Name.: Tên Volume group đang được sử dụng.

  - LV Write Access: Snapshot volume ở chế độ đọc và ghi.

  - LV Creation host, time: Thời gian khi snapshot được tạo. Nó rất quan trọng vì snapshot sẽ tìm mọi thay đổi sau thời gian này.
 
  - LV snapshot status: Snapshot này thuộc vào logical volume projects.

  - LV Size: Kích thước của Source volume mà bạn đã snapshot.
 
  - COW-table size: Kích thước bảng Cow.
 
  - Snapshot chunk size: Cung cấp kích thước của chunk cho snapshot.

### Tăng dung lượng snapshot trong LVM

- Chúng ta cần mở rộng kích thước của snapshot, chúng ta có thể thực hiện bằng cách sử dụng:

<img src="https://imgur.com/ZstxEmA.png">

- Sau khi thực thi lệnh trên thì snapshot có kích thước 2GB.

- Tiếp theo, kiểm tra kích thước mới và bảng COW bằng lệnh sau.

<img src="https://imgur.com/NPS5jcp.png">

### 3.2 Restore logical volume

- Để restore snapshot, chúng ta cần hủy gắn kết hệ thống tệp.

<img src="https://imgur.com/WHg1hCC.png">

- Kiểm tra xem điểm gắn kết còn hay không:

<img src="https://imgur.com/XnhWrpq.png">

- Điểm gắn kết của chúng ta đã được huỷ, vì vậy chúng ta có thể tiếp tục restore snapshot. Để restore snapshot bằng lệnh lvconvert.

<img src="https://imgur.com/irOPxan.png">

- Sau khi merge hoàn thành, snapshot volume sẽ tự động bị xóa.

- Chúng ta kiểm tra lại /dev/vg0/projects để xem kết quả:

<img src="https://imgur.com/TjHpIMH.png">

## Nguồn tham khảo

https://news.cloud365.vn/lvm-gioi-thieu-ve-logical-volume-manager/

https://blogd.net/linux/quan-ly-phan-vung-dia-cung-tren-linux/