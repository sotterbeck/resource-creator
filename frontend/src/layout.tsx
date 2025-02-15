import { Outlet } from "react-router";
import { NavItem } from "@/components/nav-item";
import { cn } from "@/lib/utils";

export function Layout() {
  const isMac = navigator.userAgent.includes("Mac OS X");
  return (
    <div className="flex h-screen flex-col">
      <header
        className={cn(
          "flex w-full items-center justify-end px-2",
          isMac ? "draggable h-10" : "pb-2",
        )}
      >
        <NavItem to="/">Pattern</NavItem>
        <NavItem to="/random">Random</NavItem>
      </header>
      <main className="h-full">
        <Outlet />
      </main>
    </div>
  );
}
