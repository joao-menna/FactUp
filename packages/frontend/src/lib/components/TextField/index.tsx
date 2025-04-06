import { clsx } from "clsx/lite";
import { JSX } from "react";

export function TextField({
  className,
  ...rest
}: JSX.IntrinsicElements["input"]) {
  return (
    <input
      className={clsx(
        "border-2 border-accent-500 bg-primary-800",
        "rounded-lg p-1 outline-0 text-text-100",
        className
      )}
      {...rest}
    />
  );
}
