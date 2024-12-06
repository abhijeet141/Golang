CREATE TABLE CUSTOMER(
    id INT PRIMARY KEY,
    cname VARCHAR(255),
    address VARCHAR(255),
    gender CHAR(2),
    city VARCHAR(255),
    pincode INT
)

INSERT INTO CUSTOMER (id, cname, address, gender,city,pincode) VALUES
						   (1,'Abhijeet','DLW','M','VARANASI',221109),
                           (2,"Aditya",'Bazardiha','M','RANCHI',220001),
                           (3,"Shivam",'Sigra','M','LUCKHNOW',220021),
                           (4,"Rohit",'Chaheru','M','JALANDHAR',221001);

INSERT INTO CUSTOMER VALUES (5,'Shivani','DLW','F','VARANASI',221109);

UPDATE CUSTOMER SET address = 'MUMBAI' WHERE ID = 2;

UPDATE CUSTOMER SET PINCODE = 10001;

-- DELETE 

DELETE FROM CUSTOMER WHERE ID = 2;

DELETE FROM CUSTOMER;

--REPLACE

REPLACE INTO CUSTOMER (id,cname,city) VALUES (7,'Nitin','GOA');

