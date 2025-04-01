import { DropdownMenuItemProps } from "@radix-ui/react-dropdown-menu";
import { DropdownMenu } from "radix-ui";
import { clsx } from "clsx/lite";

export function DropdownMenuItem({
  className,
  ...rest
}: DropdownMenuItemProps) {
  return (
    <DropdownMenu.Item
      className={clsx(
        "hover:bg-accent-400 duration-100 p-2 outline-0",
        "rounded-lg select-none cursor-pointer text-text-200",
        "hover:text-text-100",
        className
      )}
      {...rest}
    />
  );
}
