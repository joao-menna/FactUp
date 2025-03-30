import { useLogin } from "hooks/useLogin";
import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router";

export function LoginCallback() {
  const [searchParams] = useSearchParams();
  const { setBearerToken } = useLogin();
  const navigate = useNavigate();

  useEffect(() => {
    const token = searchParams.get("token");

    if (!token) {
      navigate("/login");
      return;
    }

    localStorage.setItem("authorization", token);
    setBearerToken(token);
    navigate("/");
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <></>;
}
