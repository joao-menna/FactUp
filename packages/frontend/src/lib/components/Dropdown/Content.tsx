import { PropsWithChildren } from "react";
import { DropdownMenu } from "radix-ui";
import { clsx } from "clsx/lite";

export function DropdownMenuContent({ children }: PropsWithChildren) {
  return (
    <DropdownMenu.Portal>
      <DropdownMenu.Content
        className={clsx("flex flex-col gap-2 mb-2 bg-accent-500 rounded-lg")}
      >
        {children}
      </DropdownMenu.Content>
    </DropdownMenu.Portal>
  );
}
