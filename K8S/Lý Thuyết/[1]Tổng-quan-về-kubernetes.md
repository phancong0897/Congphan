# Tổng quan về Kubernetes

### I. Kubernetes là cái gì

### 1.1. Khái niệm cơ bản

Kubernetes là một nền tảng mã nguồn mở, có thể mở rộng để quản lý các ứng dụng được đóng gói trên container, giúp thuận lợi trong việc cấu hình và tự động hoá việc triển khai ứng dụng mà không cần phải triển khai thủ công các ứng dụng trên mỗi máy chủ.

Hoặc Kubernetes cho phép chạy các ứng dụng phần mềm của mình trên hàng nghìn nút máy tính như thể tất cả các nút đó là một máy tính khổng lồ. Nó loại bỏ cơ sở hạ tầng cơ bản và bằng cách đó, đơn giản hóa việc phát triển, triển khai và quản lý cho cả nhóm phát triển và hoạt động. Việc triển khai các ứng dụng thông qua Kubernetes luôn đơn giản, cho dù cluster của bạn chỉ chứa một vài hoặc hàng nghìn nút trong số đó, kích thước của Cluster ảnh hưởng gì cả.

Kubernetes là một hệ sinh thái lớn và phát triển nhanh chóng. Các dịch vụ, sự hỗ trợ và công cụ có sẵn rộng rãi.

### 1.2. Lợi ích

- Kubernetes đơn giản hóa rất nhiều vấn đề phức tạp trong công tác quản lý container. Và nó thực sự đem lại rất nhiều lợi ích to lớn đối với các quản trị viên như:

    - Đơn giản hóa khai thác ứng dụng: Vì Kubernetes hiển thị tất cả các nút worker của mình như một nền tảng triển khai duy nhất, các nhà phát triển ứng dụng có thể bắt đầutriển khai các ứng dụng của riêng họ và không cần quan tâm nhiều về các thông số phần cứng phức tạp trên các máy chủ tạo nên cluster.

    - Tối ưu hóa phần cứng: Kubernetes sẽ tự động chọn nút thích hợp nhất để chạy ứng dụng của bạn dựa trên mô tả về các yêu cầu tài nguyên của ứng dụng với các tài nguyên cósẵn trên mỗi nút. Nếu triển khai thủ công bạn sẽ chủ động triển khai ứng dụng trên một máy cố định nào đó, điều này không tối ưu hóa được phần cứng trong suốt quá trìnhphát triển ứng dụng. Nhưng Kubernetes sẽ tự động phân tích mức độ yêu cầu tài nguyên của mỗi giai đoạn và cho ứng dụng chạy trên máy có lượng tài nguyên đáp ứng, vì vậy ởcác thời điểm phát triển khác nhau ứng dụng có thể chạy trên các nút khác nhau.

    - Kiểm tra hóa và khắc phục lỗi: Kubernetes giám sát các thành phần ứng dụng và các nút chạy ứng dụng trên đó và tự động xử lý cho phép ứng dụng chạy trên một nút khác nếunút đó bị lỗi. Trong thực tế để di chuyển một ứng dụng từ máy này sang máy khác là rất vất vả và mất rất nhiều thời gian. Nhưng Kubernetes thực hiện điều này một cách tựđộng và cực kỳ nhanh chóng. Thời gian downtime của ứng dụng gần như bằng 0.
    
    - Tự động hóa mở rộng: Kubernetes giám sát lưu lượng truy cập vào ứng dụng. Và có thể điều chỉnh số lượng phiên bản đang chạy của mỗi ứng dụng. Kubernetes có thể tự độngđiều chỉnh kích thước toàn bộ cluster lên hoặc xuống dựa trên nhu cầu của các ứng dụng đã triển khai. Đây là một điểm mạnh về khả năng mở rộng của Kubernetes so với Dockerswarm.
    
    - Đơn giản hóa phát triển ứng dụng: Kubernetes hỗ trợ phát hiện lỗi của ứng dụng. Điều này đồng nghĩa với việc giảm rủi ro khi đưa ra ứng dụng và giảm thời khắc phục lỗigiúp việc phát triển ứng dụng giảm bớt công việc áp lực của nhà phát triển ứng dụng, tăng độ tin tưởng của người dùng với ứng dụng và nhà phát triển. Thông thường, ứng dụngchỉ tra cứu một số biến môi trường nhất định hoặc thực hiện tra cứu DNS nhưng Kubernetes hỗ trợ ứng dụng có thể truy vấn trực tiếp máy chủ Kubernetes API để lấy một sốthông tin khác.

### II. Kiến trúc Kubernetes

### 2.1. Kiến trúc tổng quan một cluster

- Kiến trúc cluster cơ bản

<h3 align="center"><img src="../Images/1.png"></h3>

Kubernetes gom nhiều node lại thành một cluster, tối thiểu sẽ có một node cùng một control plane. Hoặc có thể hiểu là worker node và master node.
Control plane có: APIserver, control manager, etcd, scheduler, cloud control manager (tùy chọn).
Node có: kubelet, kube proxy, container runtime.

Control plane chịu trách nhiệm đưa ra hầu hết các quyết định và hoạt động như một bộ não của một cluster. Node chịu trách nhiệm về việc thực thi, chạy các ứng dụng. Chúng được quản lý bởi Control plane.

Một cluster có thể bao gồm nhiều worker node và một control plane. Một worker node có thể bao gồm nhiều service. Một service có thể bao gồm nhiều pod. Một pod có thể bao gồm nhiều container runtime. Và một control plane có thể bao gồm nhiều master node.

Ngoài ra, các Cloud provider phổ biến như AWS, Azure, Digital Ocean, Google Cloud,...

- Kiến trúc cluster đặc biệt minikube

<h3 align="center"><img src="../Images/2.png"></h3>

Kiến trúc của minikube khá đặc biệt khi toàn bộ control plane và node đều nằm trên cùng một node. Lúc này node vừa đóng vai trò làm worker node vừa đóng vai trò làm master node. Minikube chỉ chạy trên môi trường local.


### 2.2. Các thành phần trong cluster cơ bản

- Nút master cung cấp môi trường cho control plane chịu trách nhiệm quản lý trạng thái của các cụm Kubernetes và nó là đầu não đằng sau mọi quá trình vận hành trong cụm.Thành phần control plane với vai trò vô cùng đặc thù trong bộ quản lý cụm. Để có thể giao tiếp với cụm Kubernetes, người sử dụng gửi các requests đến controlplane thông qua các công cụ Command Line Interface (CLI), một Dashboard dưới dạng giao diện người dùng web (Web UI) , hoặc các API (Application Programming Interface).

- Việc giữ control plane chạy bằng mọi giá là tối quan trọng. Việc control plane ngừng hoạt động có thể dẫn đến downtime, thứ mà gây ra việc các dịch vụ gián đoạn và khôngthể cung cấp cho người dùng và có thể làm kinh doanh thua lỗ. Để đảm bảo cho khả năng chịu lỗi của control plane, các bản sao của nút master được thêm vào cụm, được cấuhình ở chế độ HA (High-Availability, độ khả dụng cao). Trong khi chỉ một trong nút master được chuyên biệt để quản lý chủ động cụm, các thành phần control plane giữ trạngthái đồng bộ giữa các bản sao của nút master. Chính loại cấu hình này đã cung cấp khả năng tự phục hồi cho control plane của cụm, thứ có thể cho phép nút master đang ởtrạng thái active xảy ra lỗi ở mức độ nào đó.

- Để duy trình trạng thái của cụm, tất cả các dữ liệu về cấu hình của cụm đều được lưu trong etcd. etcd là một bộ lưu trữ dạng key-value phân tán, thứ chỉ giữ các dữ liệuliên quan đến trạng thái của cụm chứ không bao gồm dữ liệu workload của client. etcd có thể được cấu hình trên nút master (stacked topology) hoặc trên một host được chuyênbiệt (external topology) để có thể giảm thiểu khả nằng mất mát dữ liệu được lưu trữ bằng cách tách riêng chúng khỏi các tác tử control plane khác.
Với ở dạng stacked topology, các bản sao HA của nút master đảm bảo khả năng tự phục hồi của dữ liệu lưu trữ trong etcd. Tuy nhiên, đây không giống như trường hợp dùng etcdở dạng external topology, khi mà host của etcd bắt buộc phải được tạo các nhân bản và tách rời để phục vụ cho chế đồ HA, cấu hình mà dẫn đến nhu cầu về các phần cứng bổsung.

- Một nút master có thể có những thành phần control plane như sau:

    - API Server (máy chủ API)

    Tất cả các tác vụ quản trị đều được phối hợp triển khai bởi kube-apiserver, thành phần control plane trung tâm được chạy trên nút master. API Server tiếp nhận các yếu cầu RESTful từ người dùng, tác tử hoạc các tác tử từ phía bên ngoài, sau đó xác nhận và xử lý chúng. Trong khi xử lý, API Server đọc trạng thái hiện tại của cụm Kubernetes từ bộ lưu trữ dữ liệu etcd và sau khi thực thi xong yêu cầu, trạng thái kết quả của cụm Kubernetes được lưu vào bộ lưu trữ dữ liệu dạng key-value phân tán để đảm bảo tính bền bỉ. API Server là thành phần master plane duy nhất có thể giao tiếp với bộ lưu trữ dữ liệu etcd. bao gồm cả đọc và ghi các thông tin trạng thái của cụm Kubernetes - hoạt động như một giao diện trung gian cho bất kì một tác tử control plane khác truy cập đến trạng thái của cụm.

    API Server có thể cấu hình và tùy biến một cách dễ dàng. Nó có thể mở rộng quy mô theo chiều ngang, nhưng nó cũng hỗ trợ việc bổ sung API Server phụ tùy chỉnh, một cấu hình biến API Server chính thành proxy cho tất cả API Server tùy chỉnh, thứ cấp và định tuyến tất cả các lệnh gọi RESTful đến chúng dựa trên các quy tắc được xác định tùy chỉnh.

    - Scheduler (Bộ lập lịch)

    Vai trò của kube-scheduler là giao các đối tượng workload mới, chẳng hạn như các pod cho các nút. Trong suốt quá trình lập lịch, các quyết định sẽ được tạo dựa trên trạng thái hiện tại của cụm Kubernetes và yêu cầu của các đối tượng mới. Bộ lập lịch thu thập từ bộ lưu trữ dữ liệu, thông qua API server, dữ liệu sử dụng tài nguyên của từng nút worker trong cụm. Bộ lập lịch cũng có thể nhận từ API Server các requirements của các đối tượng mới, thứ là một phần của dữ liệu cấu hình của chúng. Những requirements này có thể chứa các ràng buộc được người dùng hoặc các tác tử cài đặt, chẳng hạn như công việc lập lịch trên một nút đã được gán nhãn với căp key/value disk==ssd. Bộ lập lịch cũng nhận các yêu cầu về Chất lượng Dịch vụ (QoS) của tài khoản, data locality, affinity, anti-affinity, taints, toleration, cluster topology, ... Một khi tất cả các dữ liệu của cụm đã khả dụng, thuật toán lập lịch lọc các nút với các thuộc từ (predicates) để cô lập các nút ứng cử viên có thể đảm nhiệm công việc, những nút được đánh giá để có thể chọn một nút mà thỏa mãn tất cả các yêu cầu cho việc phục vụ một khối lượng công việc mới. Kết quả của quá trình đưa ra quyết định này là liên lạc lại với API Server, thứ mà sau đó ủy sẽ thác việc triển khai khối lượng công việc với các tác tử control plane khác.

    Bộ lập lịch có khả năng cấu hình và tùy chỉnh cực kì cao dựa trên các scheduling policies, plugins, and profiles. Tiếp đó, các bộ lập lịch tùy chỉnh bổ sung có thể cũng được hỗ trợ, do đó các dữ liệu cấu hình (objects' configuration data) nên bao gồm tên của bộ lập lịch tùy chỉnh mong muốn để có những quyết định lập lịch phù hợp với cài đặt của chúng. Nếu không có bất kỳ dữ liệu nào được cung cấp, bộ lập lịch mặc định sẽ được sử dụng.

    - Controller Managers (Trình quản lý các controller)
    
    controller managers là các thành phần của control plane trên nút master, thứ mà chạy các controller để điều tiết trạng thái của cụm Kubernetes. Các controller luôn trong qúa trình liên tục lặp lại việc quan sát để có thể chạy và so sánh trạng thái đã được miêu tả (được cung cấp bởi các objects' configuration data) và trạng thái hiện tại của nó (thu thập từ dữ liệu của etcd lưu trữ thông qua API server) TRong trường hợp xảy ra sự nhầm lẫn nào đó, hành động sửa lỗi được thực hiện trong cụm cho đến khi trạng thái hiện tại đúng với trạng thái được mô tả.

    - kube-controller-manager chạy các controllers chịu trách nhiệm phản ứng khi các nút không khả dụng, để đảm bảo số pod như mong đợi, để tạo endpoints, service accounts, và API access tokens.

    - cloud-controller-manager chạy các controllers chịu trách nhiệm tương tác với các hệ sinh thái cơ bản của bên cung cấp dịch vụ đám mây khi có nút trở nên không khả dụng cùng với đó là quản lý các container dữ liệu được cung cấp bởi các dịch vụ đám mấy và quản lý quá trình cân bằng tải và điều hướng.

    - Data Store (Kho dữ liệu)
    
    etcd là kho dữ liệu dạng key/value phân tán và có có tính nhất quán cao được sử dụng để đảm bảo tính bền bỉ của trạng thái của cụm Kubernetes. Dữ liệu mới được ghi vào kho lưu trữ chỉ bằng cách thêm vào cuối nó bởi vậy nên dữ liệu sẽ không bao giờ bị thay đổi trong đây. Dữ liệu lỗi thời được nén định kỳ để giảm thiểu dung lượng của kho dữ liệu.


    Công cụ quản lý dưới dạng CLI của etcd - etcdctl cung cấp khả năng backup, snapshot, and restore, những thứ đặc biệt tiện dụng đối với một cụm Kubernetes có một etcd duy nhất - thường thấy trong trong môi trường Phát triển. Tuy nhiên, trong môi trường Staging và Production, điều cực kỳ quan trọng là phải sao chép các kho dữ liệu ở chế độ HA, để có khả năng phục hồi dữ liệu cấu hình cụm.

Trong tất cả các thành phần của control plane, chỉ có API Server là có khả năng giao tiếp với kho dữ liệu etcd

- Nút worker

worker node cung cấp môi trường chạy cho các ứng dụng client. Trong vô vàng các microservices đã được container hóa, các ứng dụng được cô lập trong các Pods, được điều khiển bằng các tác tử của control plane của cụm chạy trên nút master. Pods được lập lịch trên nút worker, nơi chúng có thể tìm thấy các tài nguyên để tính toán và lưu trữ để chạy và được kết nối để có thể giao tiếp với nhau cũng như với thế giới bên ngoài. Một Pod là đơn vị lập lịch nhỏ nhất của Kubernetes. Nó là một tập hợp dựa trên logic của một hoặc nhiều container đã được lập lịch cùng với nhau và tập hợp đó có để được khởi chạy, dừng hoắc tái lập lịch như một đơn vị riêng biệt của công việc.

Mặt khác, trên một cụm Kubernetes có nhiều worker, lưu lượng mạng giữa các client users và các ứng dụng được triển khai trên các Pods được xử lý trực tiếp bởi các nút worker và nó không được điều hướng thông qua nút master.

- Một nút worker có các thành phần như sau:

    - Container Runtime (môi trường chạy của container)

    Mặc dù Kubernetes được thiết kế như một engine có khả năng điều phối container, nó lại không có khả năng để có thể xử lý trực tiếp các container. Để quản lý vòng đời của một container, Kubernetes cần có container runtime trên nuts mà một Pod và các container của nó được lập lịch. Kubernetes hỗ trợ các môi trường runtime sau:

        Docker - container runtime phổ biến nhất được sử dụng với Kubernetes (nhưng k8s đã ngừng hỗ trợ tại thời điểm gần đây)

        CRI-O - container runtime cho Kubernetes, nó cũng hỗ trợ Docker image registries

        containerd - container runtime đơn giản và di động cung cấp sự tự động đáng kể

        frakti - container runtime dựa trên hypervisor cho Kubernetes

    - kubelet

    kubelet là tác tử chạy trên mỗi nút và giao tiếp với các thành phần của control plane từ nút master. Nó nhận thông tin định nghĩa Pod, chủ yếu từ API Server và tương tác với container runtime ở trên nút để chạy các container liên quan đến Pod đó. Nó cũng theo dõi tình trạng và tài nguyên của các Pod đang chạy các containers.

    kubelet kết nối với các container runtime thông qua interface dựa trên plugin - Container Runtime Interface (CRI). CRI bao gồm bộ đệm protocol, gRPC API, thư viện cũng như các thông số kỹ thuật và công cụ bổ sung hiện đang được phát triển. Để kết nối với container runtime có thể hoán đổi cho nhau, kubelet sử dụng ứng dụng shim để cung cấp một lớp trừu tượng rõ ràng giữa kubelet và container runtime

    - Proxy - kube-proxy
    
    kube-proxy là tác nhân mạng chạy trên mỗi nút chịu trách nhiệm cập nhật và bảo trì động của các luật điều phối mạng trên nút đó. Nó trừu tượng hóa chi tiết của quá trình hoạt động mạng của Pod và chuyển tiếp các kết nối đến Pods.

    The kube-proxy is responsible for TCP, UDP, and SCTP stream forwarding or round-robin forwarding across a set of Pod backends, and it implements forwarding rules defined by users through Service API objects.

    kube-proxy chịu trách nhiệm cho việc chuyển tiếp dòng TCP, UDP và SCTP hoặc xoay vòng chuyển tiếp qua một tập các Pod backends và nó cũng cài đặt các luật chuyển tiếp được định nghĩa bởi người dùng thông qua các đối tượng Server API.
