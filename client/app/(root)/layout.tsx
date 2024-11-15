import type { Metadata } from "next";
import { ReactNode } from "react";
// import localFont from "next/font/local";

export const metadata: Metadata = {
  title: "Regular Billing",
};

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
