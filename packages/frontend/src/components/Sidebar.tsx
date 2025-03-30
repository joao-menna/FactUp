import { SidebarBottom } from "./SidebarBottom";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { FaPlus } from "react-icons/fa";
import { clsx } from "clsx/lite";

export function Sidebar() {
  const { t } = useTranslation();

  return (
    <div className="w-64 bg-primary-500 p-2 flex flex-col justify-between">
      <div className={clsx("flex flex-col items-center gap-2")}>
        <h1 className="text-text-100 text-2xl pt-2 pb-4 select-none">
          {t("applicationName")}
        </h1>
        <Button
          className={clsx(
            "bg-accent-400 hover:bg-accent-400/80 w-full py-4 flex",
            "items-center justify-between"
          )}
        >
          <FaPlus />
          <span className="text-xl">{t("post")}</span>
          <FaPlus />
        </Button>
        <Button
          className={clsx("bg-accent-500 hover:bg-accent-500/80 w-full py-2")}
        >
          {t("curiosities")}
        </Button>
        <Button
          className={clsx("bg-accent-500 hover:bg-accent-500/80 w-full py-2")}
        >
          {t("sayings")}
        </Button>
      </div>
      <SidebarBottom />
    </div>
  );
}
