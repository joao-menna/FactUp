import { clsx } from "clsx/lite";
import { JSX } from "react";

type ButtonProps = JSX.IntrinsicElements["button"];

interface Props extends ButtonProps {
  disabled?: boolean;
}

export function Button({ children, className, disabled, ...rest }: Props) {
  return (
    <button
      className={clsx(
        "rounded-lg px-2 py-1.5 duration-200",
        "text-text-200 hover:text-text-100",
        "hover:cursor-pointer select-none",
        disabled && "opacity-50",
        className
      )}
      {...rest}
    >
      {children}
    </button>
  );
}
