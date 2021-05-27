# Công nghệ RAID

- Trong vài năm trở lại đây, từ chỗ là một thành phần “xa xỉ” chỉ có trên các hệ thống máy tính lớn, máy trạm, máy chủ, RAID đã được đưa vào các máy tính để bàn dưới dạng tích hợp đơn giản. Tuy nhiên, có thể người mua biết bo mạch chủ (BMC) của mình có công nghệ RAID nhưng không phải ai cũng biết cách sử dụng hiệu quả. Bài viết này giới thiệu thông tin cơ bản về RAID cũng như một vài kinh nghiệm sử dụng để tăng sức mạnh cho PC.

![Imgur](https://imgur.com/kztSNKw.png)

### 1. RAID LÀ GÌ?

- RAID là chữ viết tắt của Redundant Array of Independent Disks. Ban đầu, RAID được sử dụng như một giải pháp phòng hộ vì nó cho phép ghi dữ liệu lên nhiều đĩa cứng cùng lúc. Về sau, RAID đã có nhiều biến thể cho phép không chỉ đảm bảo an toàn dữ liệu mà còn giúp gia tăng đáng kể tốc độ truy xuất dữ liệu từ đĩa cứng. Dưới đây là năm loại RAID được dùng phổ biến:

- Có thể hiểu nhanh về RAID qua thông tin dưới đây :

    - RAID chỉ nên làm việc với các loại ổ cứng dung lượng bằng nhau.

    - Sử dụng RAID sẽ tốn số lượng ổ nhiều hơn bình thường, nhưng đổi lại là dữ liệu sẽ an toàn hơn.
    
    - RAID có thể dùng cho bất kỳ hệ điều hành nào, từ Window 98, window 2000, window XP, Window 10, window server 2016, MAC OS X, Linux…vv
    
    - RAID 0 bằng tổng dung lượng các ổ cộng lại.
    
    - RAID 1 chỉ duy trì dung lượng 1 ổ.

    - RAID 5 sẽ có dung lượng ít hơn 1 ổ (5 ổ dùng raid 5 sẽ có dung lượng 4 ổ).

    - RAID 6 sẽ có dung lượng ít hơn 2 ổ (5 ổ dùng raid 6 sẽ có dung lượng 3 ổ).

    - RAID 10 sẽ chỉ tạo được khi số ổ là chẵn, phải có tối thiểu từ ô ổ trở lên. Dung lượng bằng tổng số ổ chia đôi (10 ổ thì dung lượng sử dụng là 5 ổ).

### 2. Các loại RAID

Về phân loại thì hiện nay có khá nhiều loại Raid được sử dụng như Raid 0, Raid 1, Raid 3, Raid 4, Raid 5, Raid 10 .v.v… Trong bài viết tìm hiểu về Raid ngày hôm nay chúng tôi xin được giới thiệu với các bạn top các loại Raid chính phổ biến và thường được các khách hàng sử dụng.

Chi tiết hơn, các bạn có thể xem ở phía dưới đây:

### 2.1 RAID 0

![Imgur](https://imgur.com/eheSeev.png)

Đây là dạng RAID đang được người dùng ưa thích do khả năng nâng cao hiệu suất trao đổi dữ liệu của đĩa cứng. Đòi hỏi tối thiểu hai đĩa cứng, RAID 0 cho phép máy tính ghi dữ liệu lên chúng theo một phương thức đặc biệt được gọi là Striping. Ví dụ bạn có 8 đoạn dữ liệu được đánh số từ 1 đến 8, các đoạn đánh số lẻ (1,3,5,7) sẽ được ghi lên đĩa cứng đầu tiên và các đoạn đánh số chẵn (2,4,6,8) sẽ được ghi lên đĩa thứ hai. Để đơn giản hơn, bạn có thể hình dung mình có 100MB dữ liệu và thay vì dồn 100MB vào một đĩa cứng duy nhất, RAID 0 sẽ giúp dồn 50MB vào mỗi đĩa cứng riêng giúp giảm một nửa thời gian làm việc theo lý thuyết. Từ đó bạn có thể dễ dàng suy ra nếu có 4, 8 hay nhiều đĩa cứng hơn nữa thì tốc độ sẽ càng cao hơn. Tuy nghe có vẻ hấp dẫn nhưng trên thực tế, RAID 0 vẫn ẩn chứa nguy cơ mất dữ liệu. Nguyên nhân chính lại nằm ở cách ghi thông tin xé lẻ vì như vậy dữ liệu không nằm hoàn toàn ở một đĩa cứng nào và mỗi khi cần truy xuất thông tin (ví dụ một file nào đó), máy tính sẽ phải tổng hợp từ các đĩa cứng. Nếu một đĩa cứng gặp trục trặc thì thông tin (file) đó coi như không thể đọc được và mất luôn. Thật may mắn là với công nghệ hiện đại, sản phẩm phần cứng khá bền nên những trường hợp mất dữ liệu như vậy xảy ra không nhiều.

Có thể thấy RAID 0 thực sự thích hợp cho những người dùng cần truy cập nhanh khối lượng dữ liệu lớn, ví dụ các game thủ hoặc những người chuyên làm đồ hoạ, video số.

### 2.2 RAID 1

![Imgur](https://imgur.com/hbg5P3P.png)

Đây là dạng RAID cơ bản nhất có khả năng đảm bảo an toàn dữ liệu. Cũng giống như RAID 0, RAID 1 đòi hỏi ít nhất hai đĩa cứng để làm việc. Dữ liệu được ghi vào 2 ổ giống hệt nhau (Mirroring). Trong trường hợp một ổ bị trục trặc, ổ còn lại sẽ tiếp tục hoạt động bình thường. Bạn có thể thay thế ổ đĩa bị hỏng mà không phải lo lắng đến vấn đề thông tin thất lạc. Đối với RAID 1, hiệu năng không phải là yếu tố hàng đầu nên chẳng có gì ngạc nhiên nếu nó không phải là lựa chọn số một cho những người say mê tốc độ. Tuy nhiên đối với những nhà quản trị mạng hoặc những ai phải quản lý nhiều thông tin quan trọng thì hệ thống RAID 1 là thứ không thể thiếu. Dung lượng cuối cùng của hệ thống RAID 1 bằng dung lượng của ổ đơn (hai ổ 80GB chạy RAID 1 sẽ cho hệ thống nhìn thấy duy nhất một ổ RAID 80GB).

### 2.3 RAID 0+1

Có bao giờ bạn ao ước một hệ thống lưu trữ nhanh nhẹn như RAID 0, an toàn như RAID 1 hay chưa? Chắc chắn là có và hiển nhiên ước muốn đó không chỉ của riêng bạn. Chính vì thế mà hệ thống RAID kết hợp 0+1 đã ra đời, tổng hợp ưu điểm của cả hai “đàn anh”. Tuy nhiên chi phí cho một hệ thống kiểu này khá đắt, bạn sẽ cần tối thiểu 4 đĩa cứng để chạy RAID 0+1. Dữ liệu sẽ được ghi đồng thời lên 4 đĩa cứng với 2 ổ dạng Striping tăng tốc và 2 ổ dạng Mirroring sao lưu. 4 ổ đĩa này phải giống hệt nhau và khi đưa vào hệ thống RAID 0+1, dung lượng cuối cùng sẽ bằng ½ tổng dung lượng 4 ổ, ví dụ bạn chạy 4 ổ 80GB thì lượng dữ liệu “thấy được” là (4*80)/2 = 160GB.

### 2.4 RAID 5

![Imgur](https://imgur.com/v7aQmpW.png)

Raid 5 là gì? Về cơ bản của Raid 5 cũng gần giống với 2 loại raid lưu trữ truyền thống kể trên là Raid 1 và Raid 0. Tức là cũng có thể tách ra lưu trữ các ổ cứng riêng biệt và vẫn có phương án dự phòng khi có sự cố phát sinh đối với 1 ổ cứng bất kỳ trong cụm. 

Đây có lẽ là dạng RAID mạnh mẽ nhất cho người dùng văn phòng và gia đình với 3 hoặc 5 đĩa cứng riêng biệt. Dữ liệu và bản sao lưu được chia lên tất cả các ổ cứng. Nguyên tắc này khá rối rắm. Chúng ta quay trở lại ví dụ về 8 đoạn dữ liệu (1-8) và giờ đây là 3 ổ đĩa cứng. Đoạn dữ liệu số 1 và số 2 sẽ được ghi vào ổ đĩa 1 và 2 riêng rẽ, đoạn sao lưu của chúng được ghi vào ổ cứng 3. Đoạn số 3 và 4 được ghi vào ổ 1 và 3 với đoạn sao lưu tương ứng ghi vào ổ đĩa 2. Đoạn số 5, 6 ghi vào ổ đĩa 2 và 3, còn đoạn sao lưu được ghi vào ổ đĩa 1 và sau đó trình tự này lặp lại, đoạn số 7,8 được ghi vào ổ 1, 2 và đoạn sao lưu ghi vào ổ 3 như ban đầu. Như vậy RAID 5 vừa đảm bảo tốc độ có cải thiện, vừa giữ được tính an toàn cao. Dung lượng đĩa cứng cuối cùng bằng tổng dung lượng đĩa sử dụng trừ đi một ổ. Tức là nếu bạn dùng 3 ổ 80GB thì dung lượng cuối cùng sẽ là 160GB.
