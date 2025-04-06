import { clsx } from "clsx/lite";
import { JSX } from "react";

export function CardWithoutPadding({
  className,
  ...rest
}: JSX.IntrinsicElements["div"]) {
  return (
    <div className={clsx("bg-primary-700 rounded-lg", className)} {...rest} />
  );
}
