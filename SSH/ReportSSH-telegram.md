# Cấu hình cảnh báo đăng nhạp SSH tới telegram

## Mục lục

### [1. Tạo 1 Telegram bot](https://github.com/phancong0897/Congphan/blob/master/SSH/ReportSSH-telegram.md#1-t%E1%BA%A1o-m%E1%BB%99t-telegram-bot)

### [2. Gửi cảnh báo sử dụng jq](https://github.com/phancong0897/Congphan/blob/master/SSH/ReportSSH-telegram.md#2-g%E1%BB%ADi-c%E1%BA%A3nh-b%C3%A1o-s%E1%BB%AD-d%E1%BB%A5ng-jq-1)

### [ Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/SSH/ReportSSH-telegram.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Tạo một Telegram bot

### Tạo 1 Telegram bot

- Trên thanh tìm kiếm của Telegram, tìm đến người dùng BotFather

<img src="https://imgur.com/HTmz7WK.png">

- Chat /newbot với BotFather.

    - Đặt tên cho bot của bạn. Ở đây ví dụ đặt là Congphan_SSH_bot.

    <img src="https://imgur.com/xATWwPc.png">

- Sau khi tạo bot thành công sẽ có thông báo về một token cho bot. Ở đây là: 1267912371:AAEdXOi30Wf4-Xm4qc4OiFGZ6tbCBHv9eZQ.

### Thêm bot vào group

- Ta có thể thêm bot vào bất kỳ group nào ta muốn, ở đây ta sẽ thêm vào một group riêng chuyên gửi cảnh báo SSH Report.

- Chọn Add member -> @Congphan_SSH_bot.

- Khởi động bot bằng cách chat với bot trong nhóm: Ví dụ /my_id @Congphan0897

### Lấy chat_id

- Để lấy được chat_id, ta sử dụng đường link sau:

    - ` https://api.telegram.org/bot[TOKEN]/getUpdates `

Tại ô [TOKEN] chính là Token nhận được khi tạo bot. Ví dụ là: https://api.telegram.org/1267912371:AAEdXOi30Wf4-Xm4qc4OiFGZ6tbCBHv9eZQ/getUpdates

Lúc này kết quả sẽ hiện ra, ta tìm đến đoạn "chat":{"id":-xxxxxxxxx,"title":"SSH Report". Chat_id của bạn chính là chuỗi kí tự -xxxxxxxxx.

```

{"ok":true,"result":[{"update_id":609530234,
"message":{"message_id":1,"from":{"id":856090516,"is_bot":false,"first_name":"cong","last_name":"phan","username":"congphan0897"},"chat":{"id":-xxxxxx,"title":"SSH_report","type":"group","all_members_are_administrators":true},"date":1598430683,"group_chat_created":true}},{"update_id":609530235,
"message":{"message_id":2,"from":{"id":856090516,"is_bot":false,"first_name":"cong","last_name":"phan","username":"congphan0897"},"chat":{"id":-426576502,"title":"SSH_report","type":"group","all_members_are_administrators":true},"date":1598430704,"text":"/my_id @congphan0897","entities":[{"offset":0,"length":6,"type":"bot_command"},{"offset":7,"length":13,"type":"mention"}]}}]}

```

### Kiểm tra cảnh báo tới Telegram qua API

- Ta sử dụng trên trình duyệt Web với method GET của http với cú pháp:

    - ` https://api.telegram.org/bot[TOKEN]/sendMessage?chat_id=[CHAT_ID]&text=[MY_MESSAGE_TEXT] `

    - [TOKEN]: Mã token ta nhận ở trên.

    - [CHAT_ID]: chat_id ta vừa lấy.

    - [MY_MESSAGE_TEXT]: Tin nhắn bất kỳ mà ta muốn. Ở đây ta test gửi đoạn text Xin Chao Phan Cong.

- Kết quả nhận được khi tin nhắn được chuyển tới group Telegram.

<img src="https://imgur.com/jlebHGT.png">

## 2. Gửi cảnh báo sử dụng jq

- Với cách trên, ta chỉ sử dụng để kiểm tra gửi cảnh báo về Telegram. Để thực hiện việc thông báo về đăng nhập SSH tới Telegram. Ta phải sử dụng script gửi cảnh báo để tự động hóa. Ở đây ta sẽ sử dụng ứng dụng jq

- jq là một ứng dụng để đọc thông tin file JSON trên linux. Để tìm hiểu ta có thể truy cập tại đây

- Cài đặt jq:

```
    - yum install epel-release -y

    - yum install jq -y

```

- Tạo Script: Tạo một thư mục /etc/profile.d/. Để khi đăng nhập vào hệ thống thì script sẽ thực hiện ngay lập tức.

- Tạo file script ssh-telegram.sh:

    - ` vi /etc/profile.d/ssh-telegram.sh `

- Nội dung script:

```

# ID chat Telegram
USERID="<target_user_id>"

# API Token bot
TOKEN="<bot_private_TOKEN>"

TIMEOUT="10"

# URL gửi tin nhắn của bot
URL="https://api.telegram.org/bot$TOKEN/sendMessage"

# Thời gian hệ thống
DATE_EXEC="$(date "+%d %b %Y %H:%M")"

# File temp
TMPFILE='/tmp/ipinfo.txt'

if [ -n "$SSH_CLIENT" ]; then
    IP=$(echo $SSH_CLIENT | awk '{print $1}')
    PORT=$(echo $SSH_CLIENT | awk '{print $3}')
    HOSTNAME=$(hostname -f)
    IPADDR=$(echo $SSH_CONNECTION | awk '{print $3}')

    # Lấy các thông tin từ IP người truy cập theo trang ipinfo.io
    curl http://ipinfo.io/$IP -s -o $TMPFILE
    CITY=$(cat $TMPFILE | jq '.city' | sed 's/"//g')
    REGION=$(cat $TMPFILE | jq '.region' | sed 's/"//g')
    COUNTRY=$(cat $TMPFILE | jq '.country' | sed 's/"//g')
    ORG=$(cat $TMPFILE | jq '.org' | sed 's/"//g')

    # Nội dung cảnh báo
    TEXT=$(echo -e "Thời gian: $DATE_EXEC\nUser: ${USER} logged in to $HOSTNAME($IPADDR) \nFrom $IP - $ORG - $CITY, $REGION, $COUNTRY on port $PORT")

    # Gửi cảnh báo
    curl -s -X POST --max-time $TIMEOUT $URL -d "chat_id=$USERID" -d text="$TEXT" > /dev/null

    # Xóa file temp khi script thực hiện xong
    rm $TMPFILE
fi

```

- Trong đó, ta chỉ cần điền chat_id và TOKEN mà ta nhận được ở trên sau đó lưu lại là được. 

- Lưu ý: tại chat_id phải có dấu - ở trước.

- Cấp quyền thực thi cho script:

    - chmod +x /etc/profile.d/ssh-telegram.sh

- Kiểm tra đăng nhập bằng SSH:

<img src ="https://imgur.com/Wr6sYX3.png">

## Nguồn tham khảo

https://dotrungquan.info/huong-dan-tao-bot-giam-sat-truy-cap-ssh-thong-qua-ung-dung-telegram/

https://blog.cloud365.vn/linux/tao-canh-bao-su-dung-telegram-python/