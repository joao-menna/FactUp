import { clsx } from "clsx/lite";
import { JSX } from "react";

export function Card({ className, ...rest }: JSX.IntrinsicElements["div"]) {
  return (
    <div
      className={clsx("bg-primary-700 p-2 rounded-lg", className)}
      {...rest}
    />
  );
}
