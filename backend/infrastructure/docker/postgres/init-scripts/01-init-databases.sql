-- Создаем отдельные БД для каждого сервиса
CREATE DATABASE auth_db;
CREATE DATABASE user_db;
CREATE DATABASE course_db;
CREATE DATABASE progress_db;
CREATE DATABASE quest_db;
CREATE DATABASE shop_db;
CREATE DATABASE community_db;
CREATE DATABASE notification_db;
CREATE DATABASE analytics_db;
CREATE DATABASE admin_db;

-- Создаем расширения для каждой БД
\c auth_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c user_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c course_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c progress_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c quest_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c shop_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c community_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c notification_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c analytics_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c admin_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
