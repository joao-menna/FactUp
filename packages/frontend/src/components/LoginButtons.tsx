import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useNavigate } from "react-router";
import { clsx } from "clsx/lite";

export function LoginButtons() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  const handleClickLogInButton = () => {
    navigate("/login");
  };

  return (
    <div className={clsx("flex gap-2")}>
      <Button
        className={clsx("bg-accent-500 hover:bg-accent-500/80 w-full py-2")}
        onClick={handleClickLogInButton}
      >
        {t("signUp")}
      </Button>
      <Button
        className={clsx("bg-accent-500 hover:bg-accent-500/80 w-full py-2")}
        onClick={handleClickLogInButton}
      >
        {t("logIn")}
      </Button>
    </div>
  );
}
