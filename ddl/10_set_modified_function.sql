-- レコードの更新時に modified_at に現在時刻を設定する。

CREATE FUNCTION set_modified_at() RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'UPDATE') THEN
    NEW.modified_at := now() AT TIME ZONE 'Asia/Tokyo';
    return NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;
