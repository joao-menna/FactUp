import { useTranslation } from "react-i18next";
import { clsx } from "clsx/lite";
import { Button } from "lib/components/Button";
import { useNavigate } from "react-router";

export function Page404() {
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
      <h1 className="text-2xl">{t("pageNotFound")}</h1>
      <p>{t("weHadTroubleTryingToFindThisPage")}</p>
      <Button className="bg-accent-400" onClick={handleClickReturn}>
        {t("returnToASafePlace")}
      </Button>
    </div>
  );
}
