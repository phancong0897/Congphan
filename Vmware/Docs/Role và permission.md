# Phân quyền cho hệ thống Vcenter 

- Mô hình phân quyền cho hệ thống vCenter Server dựa vào việc gán quyền cho các đối tượng trong cấu trúc phân cấp đối tượng.

- Mỗi đối tượng có các quyền liên quan, mỗi quyền chỉ định cho một group hoặc user mà group hoặc user đó có đặc quyền trên đối tượng

VD: Trên một máy ảo có thể chọn thêm quyền để gán vai trò cho user hoặc nhóm user, vai trò này cung cấp cho những user đó các đặc quyền tương ứng trên máy ảo

- Các khái niệm
    
    - Object: một thực thể mà các hành động được thực hiện như Datacenter, folder, resource pools, cluster, host, VM.

    - Permissions: quyền cho mỗi đối tượng trên vCenter

    - Privileges: đặc quyền – kiểm soát truy cập chi tiết vào tài nguyên, các đặc quyền nhóm thành các vai trò được ánh xạ tới user hoặc group

    - Users and groups: Trên hệ thống máy chủ vCenter, chỉ có thể chỉ định đặc quyền cho user hoặc nhóm user đã được xác thực thông qua đăng nhập một lần (SSO) hoặc từ các nguồn nhận dạng bên ngoài nà vCenter đang sử dụng như Microsoft Active Directory.

    - Roles: Vai trò là tập hợp các đặc quyền. Vai trò cho phép chỉ định quyền trên một đối tượng dựa trên một nhóm tác vụ điển hình mà người dùng thực hiện. Các vai trò mặc định, chẳng hạn như Administrator, được xác định trước trên vCenter và không thể thay đổi. Các vai trò khác, chẳng hạn như Resource Pool Administrator, là các vai trò mẫu được xác định trước, có thể tạo các vai trò tùy chỉnh từ đầu hoặc bằng cách sao chép và sửa đổi các vai trò mẫu.

- Quyền thường phải được định nghĩa trên cả đối tượng nguồn và đối tượng đích. Ví dụ: nếu di chuyển một máy ảo, cần có các đặc quyền trên máy ảo đó, nhưng cũng có các đặc quyền trên trung tâm dữ liệu đích.

- vCenter Server cung cấp các vai trò được xác định trước, kết hợp các bộ đặc quyền được sử dụng thường xuyên:

    - No Access	- Permanent	- Một vai trò được chỉ định cho người dùng và nhóm mới. Ngăn người dùng hoặc nhóm xem hoặc thực hiện thay đổi đối với một đối tượng.

    - Read-Only - Permanent - Một vai trò cho phép người dùng kiểm tra trạng thái của một đối tượng hoặc xem chi tiết của nó, nhưng không thực hiện thay đổi đối với nó.

    - Administrator - Permanent - Một vai trò cho phép người dùng truy cập đầy đủ vào tất cả các đối tượng trên máy chủ.Phải có ít nhất một người dùng có quyền quản trị trong VMware.

    - Virtual Machine Power User - Sample Role - Một vai trò chỉ cấp quyền truy cập của người dùng cho các máy ảo. Người dùng có thể thay đổi phần cứng ảo hoặc tạo ảnh chụp nhanh của máy ảo.

    - Virtual Machine User - Sample Role - Cấp quyền truy cập của người dùng dành riêng cho máy ảo. Người dùng có thể bật, tắt nguồn và đặt lại máy ảo cũng như chạy phương tiện từ các đĩa ảo

    - Resource Pool Administrator - Sample Role - Cho phép người dùng tạo các nhóm tài nguyên (RAM và CPU dành riêng để sử dụng) và gán các nhóm này cho các máy ảo.

    - Datastore Consumer - Sample Role - Cho phép người dùng sử dụng không gian trên một kho dữ liệu

    - Network Consumer - Sample Role - Cho phép người dùng chỉ định mạng cho một 
    máy ảo hoặc một máy chủ.

### Tạo các group user AD có các quyền

- Group user có quyền quản trị VM (tạo, thêm, sửa, xóa)

- Group user chỉ có quyền backup

- Group user chỉ được monitor

- Group user có chỉ được lấy log

- Group user quản lí Storage

- Group user quản lí Network

#### 1. Đối với nhóm người dùng có quyền quản trị VM, Role gồm các đặc quyền sau: Cung cấp quyền tương tác và cấu hình máy ảo

- Datastore:

    Browse datastore

- Global:

    Cancel task
    
    Scheduled task:
    
    Create tasks
    
    Modify task
    
    Remove task
    
    Run task
    
    Virtual machine:
    
    Change Configuration
    
    Acquire disk lease
    
    Add existing disk
    
    Add new disk
    
    Add or remove device
    
    Advanced configuration
    
    Change CPU count
    
    Change Memory
    
    Change Settings
    
    Change resource
    
    Modify device settings
    
    Remove disk
    
    Rename
    
    Reset guest information

    Upgrade virtual machine compatibility

- Edit Inventory:
    
    Create from existing
    
    Create new

- Interaction:

    Answer question

    Configure CD media
    
    Configure floppy media
    
    Connect devices
    
    Console interaction
    
    Guest operating system management by VIX API
    
    Install VMware Tools
    
    Power off
    
    Power on
    
    Reset
    
    Suspend

- Provisioning:
    
    Clone virtual machine
    
    Create template from virtual machine

- Snapshot management:
    
    Create snapshot
    
    Remove snapshot
    
    Rename snapshot
    
    Revert to snapshot

Dựa trên Role Virtual machine power user (sample)

#### 2. Đối với user có quyền để hoạt động backup:

- Cryptographic operations
    
    Add disk
    
    Direct Access
    
    Encrypt
    
    Encrypt new
    
    Migrate
    
- dvPort Group
    
    Create
    
    Delete
    
    Modify
    
- Datastore
    
    Allocate space
    
    Browse datastore
    
    Configure datastore
    
    Low-level file operations
    
    Remove file
    
- Extension
    
    Register extension
    
    Unregister extension

- Folder
    
    Create folder
    
    Delete folder

- Global
    
    Disable methods
    
    Enable methods
    
    Licenses
    
    Log event
    
    Manage custom attributes
    
    Set custom attribute
    
    Settings

- Host
    
    Configuration
    
    Advanced settings
    
    Maintenance
    
    Network configuration
    
    Query patch
    
    Storage partition configuration

- vSphere Tagging

    Assign or Unassign vSphere Tag

- Network

    Assign network
    
    Configure

- Resource
    
    Assign virtual machine to resource pool
    
    Create resource pool
    
    Migrate powered off virtual machine
    
    Migrate powered on virtual machine
    
    Remove resource pool

- Datastore cluster
    
    Configure a datastore cluster

- Profile-driven storage
    
    Profile-driven storage update
    
    Profile-driven storage view
    
    vApp
    
    Add virtual machine
    
    Assign resource pool
    
    Unregister

- Virtual Machine
    
    - Change Configuration
    
        Acquire disk lease
        
        Add existing disk
        
        Add new disk
        
        Add or remove device
    
        Advanced configuration
    
        Change Settings
        
        Change resource
    
        Configure RAW device*
    
        Extend virtual disk
        
        Modify device settings
        
        Remove disk
        
        Rename
        
        Set annotation
        
        Toggle disk change tracking

    - Edit Inventory
    
        Create
    
        Register
    
        Remove

        Unregister
        
    - Guest operations
        
        Guest operation modifications   
        
        Guest operation program execution
        
        Guest operation queries

    - Interaction
    
        Configure CD media
    
        Configure floppy media
    
        Console interaction
    
        Connect devices
    
        Guest operating system management by VIX API
    
        Power Off
    
        Power On
    
        Suspend
    
    - Provisioning
    
        Alow disk access
    
        Allow read-only disk access
    
        Allow virtual machine download
    
        Allow virtual machine files upload
    
        Mark as template
    
        Mark as virtual machine

    - Snapshot Management
        
        Create snapshot
        
        Remove snapshot
        
        Rename snapshot
        
        Revert to snapshot


#### 3. User có quyền monitor tương ứng với Role Read Only có sẵn

#### 4. User chỉ có quyền collect log

- Global   

    Diagnostics
    
    Licenses
    
    Settings

- Host
    
    - Configuration
    
        Change SNMP settings
    
        Hyperthreading
    
        Memory configuration
    
        Network configuration
        
        Power
        
        Security profile and firewall
        
        Storage partition configuration
    - Session
    
    View and stop sessions
    
    - System
    
    Anonymous *
    
    Read *
    
    View *