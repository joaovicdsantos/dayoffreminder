ALTER TABLE public.day_offs ALTER COLUMN initial_date TYPE timestamp USING initial_date::timestamp;
ALTER TABLE public.day_offs ALTER COLUMN final_date TYPE timestamp USING final_date::timestamp;