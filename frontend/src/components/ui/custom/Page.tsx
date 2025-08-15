import type { ReactNode } from "react";

type PageProps = {
  children?: ReactNode; // children is optional
};

export const Page = ({ children }: PageProps) => {
    return <div className="flex min-h-[80vh] w-full items-center justify-center p-6 md:p-10">
        {children}
    </div>
}