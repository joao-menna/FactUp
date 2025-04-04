import { useNavigate } from "react-router";
import { useEffect } from "react";

export function LoginCallback() {
  const navigate = useNavigate();

  useEffect(() => {
    navigate("/");
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <></>;
}
