-- PostgreSQL

UPDATE credentials
    SET password = crypt('new-password', gen_salt('bf'));
