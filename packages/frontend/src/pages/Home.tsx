import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useNavigate } from "react-router";
import { clsx } from "clsx/lite";

export function HomePage() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  return (
    <div
      className={clsx(
        "flex flex-col justify-center items-center gap-2",
        "p-2 size-full"
      )}
    >
      <p className="text-text-200 text-lg">{t("iWantToSee")}</p>
      <Button
        onClick={() => navigate("/facts")}
        className={clsx(
          "bg-primary-600 md:w-42 hover:bg-primary-600/80",
          "max-md:w-full h-16"
        )}
      >
        {t("facts")}
      </Button>
      <Button
        onClick={() => navigate("/sayings")}
        className={clsx(
          "bg-primary-600 md:w-42 hover:bg-primary-600/80",
          "max-md:w-full h-16"
        )}
      >
        {t("sayings")}
      </Button>
    </div>
  );
}
