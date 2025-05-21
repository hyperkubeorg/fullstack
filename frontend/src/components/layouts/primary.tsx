import type { ReactNode } from 'react';

type PrimaryLayoutProps = {
    children: ReactNode;
};

export function PrimaryLayout({ children }: PrimaryLayoutProps) {


    return (
        <>
            {children}
        </>
    );
}
