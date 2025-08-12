
import type { ReactNode } from "react";

type ButtonGroupProps = {
  children?: ReactNode; // children is optional
};

export const ButtonGroup = ({ children }: ButtonGroupProps) => {
    return <div className={'flex gap-4 justify-end w-full'}>
        {children}
    </div>
}