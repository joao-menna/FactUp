import { useBreakpoint } from "lib/hooks/useBreakpoint";
import { LuPanelLeftClose, LuPanelRightClose } from "react-icons/lu";
import { SidebarBottom } from "./SidebarBottom";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useEffect, useState } from "react";
import { FaPlus } from "react-icons/fa";
import { clsx } from "clsx/lite";
import { useNavigate } from "react-router";

export function Sidebar() {
  const [open, setOpen] = useState<boolean>(false);
  const breakpoint = useBreakpoint();
  const navigate = useNavigate();
  const { t } = useTranslation();

  useEffect(() => {
    if (breakpoint === "sm") {
      setOpen(false);
    } else {
      setOpen(true);
    }
  }, [breakpoint]);

  const handleClickRoute = (route: string) => {
    navigate(route);
    setOpen(false);
  };

  return (
    <div
      className={clsx(
        "max-md:absolute max-md:h-full flex duration-100",
        !open && "max-md:-left-64",
        open && "max-md:left-0"
      )}
    >
      <div
        className={clsx(
          "w-64 bg-primary-500 flex flex-col justify-between",
          "h-full"
        )}
      >
        <div className={clsx("flex flex-col items-center gap-2 p-2")}>
          <Button
            onClick={() => handleClickRoute("/")}
            className={clsx("bg-primary-600 hover:bg-primary-600/80 w-full")}
          >
            <h1 className="text-text-100 text-2xl py-2 select-none">
              {t("applicationName")}
            </h1>
          </Button>
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
            onClick={() => handleClickRoute("/curiosities")}
            className={clsx("bg-accent-500 hover:bg-accent-500/80 w-full py-2")}
          >
            {t("curiosities")}
          </Button>
          <Button
            onClick={() => handleClickRoute("/sayings")}
            className={clsx("bg-accent-500 hover:bg-accent-500/80 w-full py-2")}
          >
            {t("sayings")}
          </Button>
        </div>
        <SidebarBottom />
      </div>
      <Button
        onClick={() => setOpen(!open)}
        className={clsx(
          "bg-accent-500 hover:bg-accent-500/80 rounded-l-none",
          "md:hidden size-16 flex items-center justify-center"
        )}
      >
        {open && <LuPanelLeftClose className="size-8" />}
        {!open && <LuPanelRightClose className="size-8" />}
      </Button>
    </div>
  );
}
