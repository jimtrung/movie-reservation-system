DO $$ 
DECLARE r RECORD;
BEGIN
    FOR r IN (
        SELECT 
            c.oid::regclass AS table_name, 
            a.attname AS column_name
        FROM pg_class c
        JOIN pg_attribute a ON c.oid = a.attrelid
        WHERE c.relkind = 'r'                 -- Only tables (not views, indexes, etc.)
        AND c.relnamespace = 'public'::regnamespace  -- Only user tables (not system tables)
        AND a.attnum > 0 
        AND a.attname LIKE '%_id'             -- Target columns ending in "_id"
    )
    LOOP
        EXECUTE format(
            'SELECT setval(pg_get_serial_sequence(''%I'', ''%I''), 1, false)', 
            r.table_name, r.column_name
        );
    END LOOP;
END $$;
