import { HTMLMotionProps, motion } from "motion/react";
import { clsx } from "clsx/lite";

type ButtonProps = HTMLMotionProps<"button">;

interface Props extends ButtonProps {
  disabled?: boolean;
}

export function Button({ children, className, disabled, ...rest }: Props) {
  return (
    <motion.button
      initial={{ scale: 1.0 }}
      whileTap={{ scale: 0.9 }}
      className={clsx(
        "rounded-lg px-2 py-1.5 duration-100",
        "text-text-200 hover:text-text-100",
        "hover:cursor-pointer select-none",
        disabled && "opacity-50",
        className
      )}
      {...rest}
    >
      {children}
    </motion.button>
  );
}
