INSERT INTO s_user (id, created_at, created_by, updated_at, updated_by, channel, deleted_at, last_login_fail_time,
                    login_fail, password, status, tenant_id, username)
VALUES ('00000000-0000-0000-0000-000000000001', '2024-03-25 15:39:00.386422', 'initialize',
        '2024-03-25 15:39:00.386422', 'initialize', 'orca', null, null, 0,
        '{bcrypt}$2a$10$w0hzVOLhKkp05JxiIKt.Pe1bhp4jaMJW0Sw92qT7nRDJD2SB7EubG', 1, '0000000001', 'administrator');

INSERT INTO s_user_info (id, created_at, created_by, updated_at, updated_by, address, avatar, birthday, deleted_at, email,
                         gender, name, nickname, phone, user_id)
VALUES ('00000000-0000-0000-0000-000000000001', '2024-03-25 15:39:00.407428', 'initialize',
        '2024-03-25 15:39:00.407428', 'initialize', null,
        'https://bali-attachment.oss-cn-shanghai.aliyuncs.com/bali/avatar/avatar.jpeg', '2024-01-25', null,
        'pettyferlove@live.cn', 1, 'Pettyfer Alex', 'Pettyfer', '13094186549', '00000000-0000-0000-0000-000000000001');
