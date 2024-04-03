-- This file is executed when the migration is rolled back
DELETE FROM s_user_info WHERE 1 = 1;
DELETE FROM s_user_role WHERE 1 = 1;
DELETE FROM s_user WHERE 1 = 1;
DELETE FROM s_role_permission WHERE 1 = 1;
DELETE FROM s_role WHERE 1 = 1;
DELETE FROM s_menu WHERE 1 = 1;
DELETE FROM s_permission WHERE 1 = 1;