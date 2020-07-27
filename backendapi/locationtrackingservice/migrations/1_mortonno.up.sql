
DELIMITER go
DROP FUNCTION IF EXISTS f_morton;
go

CREATE FUNCTION f_morton(
p_longitude DOUBLE
, p_latitude DOUBLE
)
RETURNS BIGINT UNSIGNED
DETERMINISTIC
BEGIN
DECLARE v_bit BIGINT UNSIGNED
DEFAULT 1;
DECLARE v_shift BIGINT UNSIGNED
DEFAULT 1;
DECLARE v_morton BIGINT UNSIGNED
DEFAULT 0;
DECLARE v_latitude BIGINT UNSIGNED
DEFAULT CAST((p_latitude + 90) * 11930464 AS UNSIGNED);
DECLARE v_longitude BIGINT UNSIGNED
DEFAULT CAST((p_longitude + 180) * 11930464 AS UNSIGNED);

WHILE v_bit <= v_latitude || v_bit <= v_longitude DO
IF v_bit & v_longitude THEN
SET v_morton := v_morton | v_shift;
END IF;
SET v_shift := v_shift << 1;
IF v_bit & v_latitude THEN
SET v_morton := v_morton | v_shift;
END IF;
SET v_shift := v_shift << 1;
SET v_bit := v_bit << 1;
END WHILE;

RETURN v_morton;
END;
go

DELIMITER ;
