import { DropdownMenuContent, DropdownMenuItem } from "lib/components/Dropdown";
import { useTranslation } from "react-i18next";
import { FaChevronUp } from "react-icons/fa";
import { DropdownMenu } from "radix-ui";
import { clsx } from "clsx/lite";
import { MouseEvent } from "react";

export function ProfileDropdown() {
  const { t } = useTranslation();

  const handleClickLogOut = (e: MouseEvent<HTMLDivElement>) => {
    e.stopPropagation();

    const baseUrl = import.meta.env.VITE_BACKEND_BASE_URL ?? "";
    location.href = `${baseUrl}/api/v1/auth/logout`;
  };

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger
        className={clsx(
          "bg-accent-500 p-2 rounded-full hover:bg-accent-500/80",
          "duration-100 outline-0"
        )}
      >
        <FaChevronUp />
      </DropdownMenu.Trigger>
      <DropdownMenuContent>
        <DropdownMenuItem onClick={handleClickLogOut}>
          {t("logOut")}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu.Root>
  );
}
