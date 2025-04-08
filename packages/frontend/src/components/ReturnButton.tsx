import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useNavigate } from "react-router";
import { clsx } from "clsx/lite";

export function ReturnButton() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  const handleClick = () => {
    navigate(-1);
  };

  return (
    <Button
      className={clsx("w-full bg-accent-500 hover:bg-accent-500/80")}
      onClick={handleClick}
    >
      {t("return")}
    </Button>
  );
}
