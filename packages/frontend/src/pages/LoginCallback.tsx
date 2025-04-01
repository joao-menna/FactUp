import { useNavigate, useSearchParams } from "react-router";
import { useEffect } from "react";

export function LoginCallback() {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  useEffect(() => {
    const token = searchParams.get("token");

    if (!token) {
      navigate("/login");
      return;
    }

    navigate("/");
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <></>;
}
