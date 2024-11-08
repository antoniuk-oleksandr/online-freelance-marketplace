INSERT INTO roles(name)
VALUES ('USER'), ('ADMIN');

INSERT INTO files (name)
VALUES ('image_1.jpg'),
       ('image_2.jpg'),
       ('image_3.jpg'),
       ('image_4.jpg'),
       ('image_5.jpg'),
       ('image_6.jpg'),
       ('image_7.jpg'),
       ('image_8.jpg'),
       ('image_9.jpg'),
       ('image_10.jpg'),
       ('image_11.jpg'),
       ('image_12.jpg'),
       ('image_13.jpg'),
       ('image_14.jpg'),
       ('image_15.jpg'),
       ('image_16.jpg'),
       ('image_17.jpg'),
       ('image_18.jpg'),
       ('image_19.jpg'),
       ('image_20.jpg');



INSERT INTO users (email, first_name, level, password, surname, private_key, public_key, created_at, about, role_id, username, avatar_id)
VALUES
    ('alex@gmail.com', 'Alex', 1, 'name', 'Smith', 'private_key_1', 'public_key_1', NOW(),
     'Freelancer specializing in web development.', 1, 'name', 1),
    ('john@gmail.com', 'John', 2, 'abcdef', 'Doe', 'private_key_2', 'public_key_2', NOW(),
     'Experienced graphic designer with a passion for creativity.', 1, 'johndoe', 2),
    ('jane@gmail.com', 'Jane', 1, 'qwerty', 'Doe', 'private_key_3', 'public_key_3', NOW(),
     'Content writer with expertise in SEO and digital marketing.', 1, 'janedoe', 3),
    ('michael@gmail.com', 'Michael', 3, 'password123', 'Johnson', 'private_key_4', 'public_key_4', NOW(),
     'Senior developer with extensive experience in mobile app development.', 1, 'michaelj', 4),
    ('emily@gmail.com', 'Emily', 2, 'letmein', 'Davis', 'private_key_5', 'public_key_5', NOW(),
     'Professional photographer with a keen eye for detail.', 1, 'emilyd', 5),
    ('chris@gmail.com', 'Chris', 1, 'pass1234', 'Brown', 'private_key_6', 'public_key_6', NOW(),
     'UI/UX designer focused on user-centric design solutions.', 1, 'chrisbrown', 6),
    ('sarah@gmail.com', 'Sarah', 2, 'adminpass', 'Miller', 'private_key_7', 'public_key_7', NOW(),
     'Digital marketer with a track record of successful campaigns.', 1, 'sarahm', 7),
    ('david@gmail.com', 'David', 3, 'mypassword', 'Wilson', 'private_key_8', 'public_key_8', NOW(),
     'Business consultant with expertise in project management.', 1, 'davidw', 8),
    ('lisa@gmail.com', 'Lisa', 1, 'userpass', 'Moore', 'private_key_9', 'public_key_9', NOW(),
     'Creative writer with a focus on engaging storytelling.', 1, 'lisam', null),
    ('james@gmail.com', 'James', 2, 'hunter2', 'Taylor', 'private_key_10', 'public_key_10', NOW(),
     'SEO specialist with a passion for improving search rankings.', 1, 'jamest', null),
    ('karen@gmail.com', 'Karen', 1, 'secret123', 'Anderson', 'private_key_11', 'public_key_11', NOW(),
     'Web developer skilled in front-end and back-end technologies.', 1, 'karenanderson', null),
    ('robert@gmail.com', 'Robert', 3, 'securepass', 'Thomas', 'private_key_12', 'public_key_12', NOW(),
     'Experienced IT support specialist with a strong problem-solving ability.', 1, 'robertthomas', null),
    ('nancy@gmail.com', 'Nancy', 2, 'simplepass', 'Jackson', 'private_key_13', 'public_key_13', NOW(),
     'Virtual assistant offering administrative and organizational support.', 1, 'nancyj', null),
    ('kevin@gmail.com', 'Kevin', 1, 'mypwd123', 'White', 'private_key_14', 'public_key_14', NOW(),
     'Skilled in data analysis and business intelligence.', 1, 'kevinw', null),
    ('amy@gmail.com', 'Amy', 2, 'changeme', 'Harris', 'private_key_15', 'public_key_15', NOW(),
     'Creative graphic designer with a portfolio of impressive projects.', 1, 'amyh', null),
    ('steve@gmail.com', 'Steve', 3, 'bestpass', 'Martin', 'private_key_16', 'public_key_16', NOW(),
     'Seasoned software developer with expertise in multiple programming languages.', 1, 'stevem', null),
    ('emma@gmail.com', 'Emma', 1, 'easy123', 'Lee', 'private_key_17', 'public_key_17', NOW(),
     'Content creator with experience in blogging and video production.', 1, 'emmal', null),
    ('daniel@gmail.com', 'Daniel', 2, 'godpass', 'Walker', 'private_key_18', 'public_key_18', NOW(),
     'Cybersecurity expert specializing in risk assessment and mitigation.', 1, 'danielw', null),
    ('olivia@gmail.com', 'Olivia', 3, 'topsecret', 'Hall', 'private_key_19', 'public_key_19', NOW(),
     'Cloud computing specialist with experience in managing scalable systems.', 1, 'oliviah', null),
    ('mark@gmail.com', 'Mark', 1, 'bestpwd', 'Allen', 'private_key_20', 'public_key_20', NOW(),
     'Project manager with a proven track record of successful project delivery.', 1, 'marka', null);


INSERT INTO categories (name)
VALUES ('Web Development'),
       ('Graphic Design'),
       ('Digital Marketing'),
       ('SEO Services'),
       ('Content Writing'),
       ('Mobile App Development'),
       ('Data Analysis'),
       ('Cybersecurity'),
       ('Cloud Computing'),
       ('E-commerce Development'),
       ('UI/UX Design'),
       ('Video Production'),
       ('Photography'),
       ('Virtual Assistance'),
       ('Software Development'),
       ('Business Consulting'),
       ('IT Support'),
       ('Project Management'),
       ('Copywriting'),
       ('Social Media Management');

INSERT INTO services (title, description, category_id, freelancer_id, created_at)
VALUES ('Custom Website Development', 'Professional web development services tailored to your needs.', 1, 1, NOW()),
       ('Responsive Web Design', 'Modern, mobile-friendly web designs.', 1, 1, NOW()),
       ('E-commerce Web Development', 'Build a powerful online store with our e-commerce solutions.', 1, 1, NOW()),
       ('Logo Design', 'Unique logo design services to represent your brand.', 2, 1, NOW()),
       ('Brand Identity Design', 'Complete brand identity design for businesses.', 2, 1, NOW()),
       ('Graphic Design for Print', 'Professional designs for brochures, flyers, and more.', 2, 1, NOW()),
       ('SEO Optimization', 'Get your website to rank higher on search engines.', 4, 1, NOW());


INSERT INTO statuses (name)
VALUES ('New'),
       ('In Progress'),
       ('Completed'),
       ('Cancelled');

INSERT INTO packages (delivery_days, description, price, title)
VALUES
-- Web Development Packages
(3, 'Basic web development package for small websites', 49.99, 'Basic Web Development'),
(7, 'Advanced web development package for mid-sized projects', 149.99, 'Advanced Web Development'),
(10, 'Enterprise web development package for large-scale projects', 499.99, 'Enterprise Web Development'),

-- Mobile App Development Packages
(5, 'Basic mobile app development for simple apps', 99.99, 'Basic Mobile App Development'),
(8, 'Standard mobile app development for feature-rich apps', 199.99, 'Standard Mobile App Development'),
(12, 'Premium mobile app development for complex apps', 399.99, 'Premium Mobile App Development'),

-- Graphic Design Packages
(2, 'Basic graphic design for logos and small projects', 29.99, 'Basic Graphic Design'),
(4, 'Standard graphic design for branding and promotional materials', 79.99, 'Standard Graphic Design'),
(6, 'Premium graphic design for large branding projects', 129.99, 'Premium Graphic Design'),

-- SEO Optimization Packages
(3, 'Basic SEO package for local businesses', 99.99, 'Basic SEO Optimization'),
(6, 'Advanced SEO package for national reach', 199.99, 'Advanced SEO Optimization'),
(9, 'Enterprise SEO package for large businesses', 399.99, 'Enterprise SEO Optimization'),

-- Digital Marketing Packages
(4, 'Basic digital marketing package for social media', 59.99, 'Basic Digital Marketing'),
(8, 'Standard digital marketing package with PPC ads', 149.99, 'Standard Digital Marketing'),
(12, 'Premium digital marketing package for full online campaigns', 299.99, 'Premium Digital Marketing'),

-- Content Writing Packages
(2, 'Basic content writing package for blog posts', 49.99, 'Basic Content Writing'),
(5, 'Standard content writing package for website content', 99.99, 'Standard Content Writing'),
(10, 'Premium content writing package for technical writing', 199.99, 'Premium Content Writing'),

-- Data Analysis Packages
(5, 'Basic data analysis for small datasets', 149.99, 'Basic Data Analysis'),
(10, 'Advanced data analysis for mid-sized datasets', 299.99, 'Advanced Data Analysis'),
(15, 'Enterprise data analysis for large datasets', 599.99, 'Enterprise Data Analysis');

INSERT INTO languages (name)
VALUES ('English'),
       ('Mandarin Chinese'),
       ('Hindi'),
       ('Spanish'),
       ('French'),
       ('Arabic'),
       ('Bengali'),
       ('Ukrainian'),
       ('Portuguese'),
       ('Indonesian'),
       ('Urdu'),
       ('German'),
       ('Japanese'),
       ('Swahili'),
       ('Marathi'),
       ('Telugu'),
       ('Turkish'),
       ('Tamil'),
       ('Vietnamese'),
       ('Italian'),
       ('Korean');

INSERT INTO skills (name)
VALUES
    -- Web Development
    ('HTML & CSS'),
    ('JavaScript'),
    ('React.js'),

    -- Graphic Design
    ('Adobe Photoshop'),
    ('Illustrator'),
    ('Sketch'),

    -- Digital Marketing
    ('SEO'),
    ('Google Ads'),
    ('Content Strategy'),

    -- SEO Services
    ('On-page SEO'),
    ('Link Building'),
    ('SEO Audits'),

    -- Content Writing
    ('Blog Writing'),
    ('Copywriting'),
    ('Technical Writing'),

    -- Mobile App Development
    ('Android Development'),
    ('iOS Development'),
    ('React Native'),

    -- Data Analysis
    ('Python for Data Analysis'),
    ('SQL'),
    ('Power BI'),

    -- Cybersecurity
    ('Network Security'),
    ('Penetration Testing'),
    ('Cryptography'),

    -- Cloud Computing
    ('AWS'),
    ('Microsoft Azure'),
    ('Google Cloud'),

    -- E-commerce Development
    ('Shopify Development'),
    ('WooCommerce'),
    ('Magento'),

    -- UI/UX Design
    ('Wireframing'),
    ('User Research'),
    ('Prototyping'),

    -- Video Production
    ('Video Editing'),
    ('Motion Graphics'),
    ('Animation'),

    -- Photography
    ('Portrait Photography'),
    ('Photo Editing'),
    ('Product Photography'),

    -- Virtual Assistance
    ('Email Management'),
    ('Calendar Management'),
    ('Data Entry'),

    -- Software Development
    ('Java Programming'),
    ('C++ Development'),
    ('Python Development'),

    -- Business Consulting
    ('Market Analysis'),
    ('Business Strategy'),
    ('Financial Planning'),

    -- IT Support
    ('Help Desk Support'),
    ('Technical Troubleshooting'),
    ('System Administration'),

    -- Project Management
    ('Agile Methodologies'),
    ('Scrum Mastery'),
    ('Risk Management'),

    -- Copywriting
    ('Sales Copywriting'),
    ('Email Copywriting'),
    ('Website Copywriting'),

    -- Social Media Management
    ('Social Media Strategy'),
    ('Facebook Advertising'),
    ('Instagram Marketing');

INSERT INTO users_skills (user_id, skill_id)
VALUES
    -- User 1 (3 skills)
    (1, 1),
    (1, 2),
    (1, 3),
    -- User 2 (3 skills)
    (2, 4),
    (2, 5),
    (2, 6),
    -- User 3 (3 skills)
    (3, 7),
    (3, 8),
    (3, 9),
    -- User 4 (3 skills)
    (4, 10),
    (4, 11),
    (4, 12),
    -- User 5 (3 skills)
    (5, 13),
    (5, 14),
    (5, 15),
    -- User 6 (3 skills)
    (6, 16),
    (6, 17),
    (6, 18),
    -- User 7 (3 skills)
    (7, 19),
    (7, 20),
    (7, 21),
    -- User 8 (3 skills)
    (8, 22),
    (8, 23),
    (8, 24),
    -- User 9 (3 skills)
    (9, 25),
    (9, 26),
    (9, 27),
    -- User 10 (3 skills)
    (10, 28),
    (10, 29),
    (10, 30),
    -- User 11 (3 skills)
    (11, 31),
    (11, 32),
    (11, 33),
    -- User 12 (3 skills)
    (12, 34),
    (12, 35),
    (12, 36),
    -- User 13 (3 skills)
    (13, 37),
    (13, 38),
    (13, 39),
    -- User 14 (3 skills)
    (14, 40),
    (14, 41),
    (14, 42),
    -- User 15 (3 skills)
    (15, 43),
    (15, 44),
    (15, 45),
    -- User 16 (3 skills)
    (16, 46),
    (16, 47),
    (16, 48),
    -- User 17 (3 skills)
    (17, 49),
    (17, 50),
    (17, 51),
    -- User 18 (3 skills)
    (18, 52),
    (18, 53),
    (18, 54),
    -- User 19 (3 skills)
    (19, 55),
    (19, 56),
    (19, 57),
    -- User 20 (3 skills)
    (20, 58),
    (20, 59),
    (20, 60);

INSERT INTO services_packages (service_id, package_id)
VALUES
    -- Service 1 (3 packages)
    (1, 1),
    (1, 2),
    (1, 3),
    -- Service 2 (3 packages)
    (2, 4),
    (2, 5),
    (2, 6),
    -- Service 3 (3 packages)
    (3, 7),
    (3, 8),
    (3, 9),
    -- Service 4 (3 packages)
    (4, 10),
    (4, 11),
    (4, 12),
    -- Service 5 (3 packages)
    (5, 13),
    (5, 14),
    (5, 15),
    -- Service 6 (3 packages)
    (6, 16),
    (6, 17),
    (6, 18),
    -- Service 7 (3 packages)
    (7, 19),
    (7, 20),
    (7, 21);

INSERT INTO users_languages (user_id, language_id)
VALUES
    -- Service 1
    (1, 1),
    (1, 2),
    (1, 3),
    -- Service 2
    (2, 4),
    (2, 5),
    (2, 6),
    -- Service 3
    (3, 7),
    (3, 8),
    (3, 9),
    -- Service 4
    (4, 10),
    (4, 11),
    (4, 12),
    -- Service 5
    (5, 13),
    (5, 14),
    (5, 15),
    -- Service 6
    (6, 16),
    (6, 17),
    (6, 18),
    -- Service 7
    (7, 19),
    (7, 20),
    (7, 21);

INSERT INTO reviews (rating, content)
VALUES (5, 'Outstanding service. Highly recommended!'),
       (4, 'Very good service with minor delays.'),
       (3, 'Satisfactory service, but could be improved.'),
       (2, 'Not happy with the service. Needs improvement.'),
       (1, 'Very poor service. Would not recommend.'),
       (4, 'Good experience overall. Will use again.'),
       (3, 'Service was okay, but had some issues.');

INSERT INTO orders (created_at, ended_at, customer_id, freelancer_id, review_id, service_id, service_package_id,
                    status_id)
VALUES
-- Order 1: New order (status is 'New', so ended_at is NULL)
(NOW(), NULL, 1, 1, NULL, 1, 1, 1),

-- Order 2: In Progress (status is 'In Progress', so ended_at is NULL)
(NOW(), NULL, 2, 1, NULL, 2, 2, 2),

-- Order 3: Completed (status is 'Completed', so ended_at is set to a past date, review ID: 1)
(NOW() - INTERVAL '7 days', NOW() - INTERVAL '6 days', 3, 1, 1, 3, 3, 3),

-- Order 4: Cancelled (status is 'Cancelled', so ended_at is set to a past date, review ID: 4)
(NOW() - INTERVAL '10 days', NOW() - INTERVAL '9 days', 4, 1, 4, 4, 4, 4),

-- Order 5: New order (status is 'New', so ended_at is NULL)
(NOW(), NULL, 5, 1, NULL, 5, 5, 1),

-- Order 6: Completed (status is 'Completed', so ended_at is set to a past date, review ID: 2)
(NOW() - INTERVAL '15 days', NOW() - INTERVAL '14 days', 6, 1, 2, 6, 6, 3),

-- Order 7: In Progress (status is 'In Progress', so ended_at is NULL)
(NOW(), NULL, 7, 1, NULL, 7, 7, 2),

-- Order 8: Cancelled (status is 'Cancelled', so ended_at is set to a past date, review ID: 5)
(NOW() - INTERVAL '20 days', NOW() - INTERVAL '19 days', 8, 1, 5, 1, 8, 4);

