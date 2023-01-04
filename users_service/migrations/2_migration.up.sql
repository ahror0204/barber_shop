CREATE TABLE IF NOT EXISTS permissions(
    id SERIAL PRIMARY KEY,
    user_type VARCHAR CHECK ("user_type" IN('superadmin', 'user', 'staff')) NOT NULL,
    resource VARCHAR NOT NULL,
    action VARCHAR NOT NULL,
    UNIQUE(user_type, resource, action)   
);

INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'customer', 'create');
INSERT INTO permissions(user_type, resource, action) VALUES ('user', 'customer', 'create');

INSERT INTO permissions(user_type, resource, action) VALUES ('user', 'customer', 'update');
INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'customer', 'update');

INSERT INTO permissions(user_type, resource, action) VALUES ('user', 'customer', 'delete');
INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'customer', 'delete');

INSERT INTO permissions(user_type, resource, action) VALUES ('user', 'customer', 'getme');
INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'customer', 'getme');

INSERT INTO permissions(user_type, resource, action) VALUES ('user', 'customer', 'update-password');

INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'staff', 'create');
INSERT INTO permissions(user_type, resource, action) VALUES ('staff', 'staff', 'create');

INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'staff', 'update');
INSERT INTO permissions(user_type, resource, action) VALUES ('staff', 'staff', 'update');

INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'staff', 'delete');
INSERT INTO permissions(user_type, resource, action) VALUES ('staff', 'staff', 'delete');

INSERT INTO permissions(user_type, resource, action) VALUES ('staff', 'staff', 'update-password');

INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'salon', 'create');

INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'salon', 'update');

INSERT INTO permissions(user_type, resource, action) VALUES ('superadmin', 'salon', 'delete');

