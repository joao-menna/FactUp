import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router";
import { clsx } from "clsx/lite";
import { Button } from "lib/components/Button";

export function ErrorPage() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  const handleClickReturn = () => {
    navigate("/");
  };

  return (
    <div
      className={clsx(
        "min-h-full flex flex-col items-center justify-center gap-4",
        "text-text-100 bg-primary-900 select-none"
      )}
    >
      <h1 className="text-2xl">{t("anInternalErrorHasOccurred")}</h1>
      <Button className="bg-accent-400" onClick={handleClickReturn}>
        {t("returnToASafePlace")}
      </Button>
    </div>
  );
}
