import { useBreakpoint } from "lib/hooks/useBreakpoint";
import { LuPanelLeftClose, LuPanelRightClose } from "react-icons/lu";
import { SidebarBottom } from "./SidebarBottom";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { useEffect, useState } from "react";
import { FaPlus } from "react-icons/fa";
import { clsx } from "clsx/lite";
import { useNavigate } from "react-router";
import { useQueryClient } from "@tanstack/react-query";
import { CURRENT, USER } from "constants/queryKeys";

export function Sidebar() {
  const [open, setOpen] = useState<boolean>(false);
  const queryClient = useQueryClient();
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

  const handleClickPost = () => {
    const user = queryClient.getQueryData([USER, CURRENT]);

    setOpen(false);

    if (!user) {
      navigate("/login");
      return;
    }

    navigate("/post");
  };

  return (
    <div
      className={clsx(
        "max-md:absolute flex duration-100 h-full sticky",
        "top-0 bottom-0",
        !open && "max-md:-left-64 max-md:pointer-events-none",
        open && "max-md:left-0"
      )}
    >
      <div
        className={clsx(
          "w-64 bg-primary-500 flex flex-col justify-between",
          "h-full z-20"
        )}
      >
        <div className={clsx("flex flex-col items-center gap-2 p-2")}>
          <Button
            onClick={() => handleClickRoute("/")}
            className={clsx(
              "bg-primary-600 hover:bg-primary-600/80 w-full flex justify-center"
            )}
          >
            {/* <h1 className="text-text-100 text-2xl py-2 select-none">
              {t("applicationName")}
            </h1> */}
            <img
              className={clsx("p-2")}
              src="/logo.svg"
              alt={t("applicationName")}
            />
          </Button>
          <Button
            onClick={handleClickPost}
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
            onClick={() => handleClickRoute("/facts")}
            className={clsx("bg-accent-500 hover:bg-accent-500/80 w-full py-2")}
          >
            {t("facts")}
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
      {breakpoint === "sm" && (
        <>
          <Button
            whileTap={{ scale: 1.0 }}
            onClick={() => setOpen(!open)}
            className={clsx(
              "bg-primary-500 rounded-l-none pointer-events-auto",
              "md:hidden size-14 flex items-center justify-center z-20"
            )}
          >
            {open && <LuPanelLeftClose className="size-8" />}
            {!open && <LuPanelRightClose className="size-8" />}
          </Button>
          <div
            onClick={() => setOpen(false)}
            className={clsx(
              "size-full fixed inset-0 duration-100 z-10",
              open && "bg-accent-500/25",
              !open && "bg-accent-500/0 pointer-events-none"
            )}
          />
        </>
      )}
    </div>
  );
}
