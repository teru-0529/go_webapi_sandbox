-- レコードの作成時に created_at/modified_at に現在時刻を設定する。

CREATE FUNCTION set_created_at() RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'INSERT') THEN
    NEW.created_at := now() AT TIME ZONE 'Asia/Tokyo';
    NEW.modified_at := now() AT TIME ZONE 'Asia/Tokyo';
    return NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
