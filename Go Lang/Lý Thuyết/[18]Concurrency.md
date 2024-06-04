# Giới thiệu về concurrency

- Go là một ngôn ngữ lập trình đồng thời (concurrent language), không phải ngôn ngữ lập trình song song (parallel language). Trước khi bàn về tính đồng thời - (concurrency) trong Go, chúng ta cần hiểu concurrency là gì và nó khác biệt ra sao so với tính song song (parallelism).

### Concurrency là gì?

- Concurrency là khả năng thực hiện nhiều tác vụ cùng một lúc. Để dễ hiểu, chúng ta cùng bắt đầu với một ví dụ.

- Giả sử ta thấy có một người đang chạy bộ buổi sáng. Trong quá trình chạy, dây giày của người này dần dần bị nới lỏng. Anh ta dừng lại, buộc chặt dây giày và tiếp tục chạy. Đây là một ví dụ đơn giản về concurrency: Một người có thể chạy và buộc dây giày, là người có thể thực hiện nhiều tác vụ với nhau một cách xen kẽ :)

### Parallelism là gì và khác biệt so với concurrency?

- Parallelism là việc thực hiện nhiều tác vụ tại cùng một thời điểm. Nghe có vẻ giống concurrency, nhưng nó thực sự khác biệt.

- Chúng ta cùng xem ví dụ về người chạy bộ. Giả thiết rằng, anh ta vừa chạy bộ, vừa nghe nhạc trong chiếc iPod. Trong trường hợp này, anh ta chạy bộ và nghe nhạc cùng thời điểm với nhau, tức là anh ta đang thực hiện nhiều tác vụ tại một thời điểm. Đây được gọi là parallelism.

### Concurrency và Parallelism dưới góc nhìn kỹ thuật

- Chúng ta đã hiểu về concurrency và sự khác biệt của nó với parallelism qua ví dụ trong thực tế. Bây giờ chúng ta sẽ tìm hiểu hai khái niệm này dưới góc nhìn kỹ thuật, vì chúng ta là dân kỹ thuật mà :)

- Giả sử chúng ta đang lập trình một trình duyệt web có nhiều thành phần khác nhau. Hai trong số đó là phần hiển thị trang web và thành phần tải về để tải file từ internet. Giả định chúng ta đã sắp xếp các đoạn mã trình duyệt theo cách mà mỗi thành phần có thể thực thi độc lập với nhau (điều này có thể thực hiện bằng cách sử dụng luồng - threads trong ngôn ngữ lập trình, ví dụ với Java và Go ta có thể dùng Goroutines, sẽ đề cập trong các bài tiếp theo). Khi trình duyệt chạy trong bộ vi xử lý (BVXL) đơn nhân, BVXL sẽ chuyển ngữ cảnh giữa hai thành phần của trình duyệt. Đôi lúc nó xử lý việc tải file, sau đó nó lại chuyển qua xử lý việc yêu cầu hiển thị trang web của người dùng. Đây chính là concurrency. Các tiến trình đồng thời (concurrent processes) bắt đầu tại một thời điểm khác, và chu trình thực hiện của chúng "ghi đè lên nhau" (cycle overlap). Trong trường hợp này, việc tải file và hiển thị trang web bắt đầu ở những thời điểm khác nhau và quá trình thực thi của chúng "ghi đè lên nhau".

- Thử xét cùng trình duyệt như trên, nhưng chạy trên BVXL đa nhân. Lúc này, thành phần tải file và hiển thị trang web có thể chạy đồng thời trên những nhân khác nhau. Đây chính là parallelism.

- Parallelism không phải luôn luôn cho kết quả thực thi nhanh hơn. Điều này là vì những thành phần chạy song song (parallel) phải kết nối với những thành phần khác. Ví dụ, trong trường hợp trình duyệt nói trên, khi quá trình tải file hoàn tất, nó sẽ giao tiếp và hiện ra thông báo cho người dùng. Quá trình giao tiếp này xảy ra giữa những thành phần chịu trách nhiệm tải file và những thành phần chịu trách nhiệm hiển thị giao diện người dùng. Trong những hệ thống chạy đồng thời (concurrent systems), quá trình giao tiếp này xảy ra với hao phí ít hơn; khi những thành phần chạy song song trên hệ thống đa nhân, hao phí trong quá trình giao tiếp thường cao hơn. Do vậy, những chương trình chạy song song không phải luôn luôn mang lại kết quả thực thi nhanh hơn!

### Hỗ trợ concurrency trong Go

- Concurrency là một thành phần kế thừa của ngôn ngữ lập trình Go. Concurrency được xử lý trong Go bằng cách sử dụng các Goroutine và các kênh (channels). Chúng ta sẽ bàn về nội dung này chi tiết hơn trong những bài hướng dẫn sau.
