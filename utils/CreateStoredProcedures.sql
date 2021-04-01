# Required configuration for recursive procedures
SET @@GLOBAL.max_sp_recursion_depth = 255;
SET @@session.max_sp_recursion_depth = 255;

DROP PROCEDURE IF EXISTS GetParamsString;
DELIMITER $$
CREATE PROCEDURE GetParamsString (
	IN  controlid VARCHAR(20),
   	INOUT result VARCHAR(2000)
)
BEGIN
	DECLARE finished INTEGER DEFAULT 0;
    DECLARE currParamId VARCHAR(20);
    DECLARE firstParam BOOL DEFAULT TRUE;
    DECLARE currLabel VARCHAR(100);
    DECLARE currControlParam
		CURSOR FOR 
			SELECT controls_params.paramid
            FROM controls_params
            WHERE controls_params.controlid = controlid;
	
     DECLARE CONTINUE HANDLER 
        FOR NOT FOUND SET finished = 1;
        
	SET result = '"parameters": [';    
    OPEN currControlParam;
	getControlParam: LOOP
		FETCH currControlParam INTO currParamId;
		IF finished = 1 THEN 
			LEAVE getControlParam;
		END IF;
        IF !firstParam THEN
			SET result = CONCAT(result, ', ');
        END IF;
        SET firstParam = FALSE;
		SELECT label from params where paramid = currParamId into currLabel;
        SET result = CONCAT(result, '{"id": "', currParamId, '", "label": "', currLabel, '"}');
	END LOOP getControlParam;
    SET result = CONCAT(result, '], ');
END$$
DELIMITER ;

DROP PROCEDURE IF EXISTS GetPartTree;
DELIMITER $$
CREATE PROCEDURE GetPartTree (
	IN  parentid VARCHAR(20),
    IN  childid VARCHAR(20),
    INOUT result TEXT
)
BEGIN
	DECLARE numOfChildren INT;
    DECLARE finished INTEGER DEFAULT 0; 
    DECLARE currPartId VARCHAR(20);
    DECLARE recursionResult TEXT;
    DECLARE firstRun BOOL DEFAULT TRUE;
    DECLARE childProse VARCHAR(10000);
    DECLARE childName VARCHAR(20);
	
    # Cursor for iterating through the children of a part
    DECLARE currPartChild
		CURSOR FOR 
			SELECT parts_parts.child_partid 
            FROM parts_parts
            WHERE parts_parts.parent_partid = childid;         
	DECLARE CONTINUE HANDLER 
        FOR NOT FOUND SET finished = 1;

   # Add the part's id, name, and prose information to the result.
   # Open the Json object's parts array
   SET result = CONCAT('{"id": "', childid, '", ');
   SELECT name FROM parts WHERE partid = childid INTO childName;
   SET result = CONCAT(result, '"name": "', childName, '", ');
   SELECT prose FROM parts WHERE partid = childid INTO childProse;
   SET result = CONCAT(result, '"prose": "', childProse, '", ');
   SET result = CONCAT(result, '"parts": [');   
   
   # Iterate through each of the part's children
   OPEN currPartChild;
   getPartChild: LOOP
		FETCH currPartChild INTO currPartId;
		IF finished = 1 THEN 
			LEAVE getPartChild;
		END IF;
		IF !firstRun THEN
			SET result = CONCAT(result, ", ");
		END IF;
        SET firstRun = FALSE;
        SET recursionResult = "";
        CALL GetPartTree(childid, currPartId, recursionResult);
        SET result = CONCAT(result, recursionResult);            
	END LOOP getPartChild;
	CLOSE currPartChild;        
    SET result = CONCAT(result, "]}");    

END$$
DELIMITER ;

DROP PROCEDURE IF EXISTS GetControlTree;
DELIMITER $$
CREATE PROCEDURE GetControlTree (
	IN  controlid VARCHAR(20),
    INOUT result TEXT
   	#INOUT result VARCHAR(65535)
)
BEGIN
	DECLARE numOfChildren INT;
    DECLARE paramFinished INTEGER DEFAULT 0; 
    DECLARE finished INTEGER DEFAULT 0;
    DECLARE currPartId VARCHAR(20);
    DECLARE partTree VARCHAR(20000) DEFAULT "";
    DECLARE firstRun BOOL DEFAULT TRUE;
    DECLARE paramsString VARCHAR(2000);
    
   	DECLARE currControlChild
		CURSOR FOR 
			SELECT controls_parts.partid
            FROM controls_parts
            WHERE controls_parts.controlid = controlid;         
	DECLARE CONTINUE HANDLER 
        FOR NOT FOUND SET finished = 1;	
	
    # Set the control's id and parameter value in the 
    # resulting JSON object and open it's parts array
    SET result = CONCAT('{ "id": "', controlid, '", ');    
    CALL GetParamsString(controlid, paramsString);
    SET result = CONCAT(result, paramsString);
    SET result = CONCAT(result, '"parts": [');
    
	OPEN currControlChild;
	getControlChild: LOOP
		FETCH currControlChild INTO currPartId;
		IF finished = 1 THEN 
			LEAVE getControlChild;
		END IF;
        IF !firstRun THEN
			SET result = CONCAT(result, ", ");
		END IF;
        SET firstRun = FALSE;
		SET partTree = "";
		CALL GetPartTree(controlid, currPartId, partTree);
		SET result = CONCAT(result, partTree);		         
	END LOOP getControlChild;
	CLOSE currControlChild;
    SET result = CONCAT(result, "]}");
    SELECT result;
END$$
DELIMITER ;

#call GetControlTree('ac-2', @result);
call GetPartTree('ac', 'ac-2_smt', @result);
select @result;




