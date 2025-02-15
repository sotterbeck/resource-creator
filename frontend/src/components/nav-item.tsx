import { NavLink } from "react-router";
import { cn } from "@/lib/utils";

interface NavItemProps {
  to: string;
  children: React.ReactNode;
}

export function NavItem({ to, children }: NavItemProps) {
  return (
    <NavLink
      to={to}
      end={to === "/"}
      className={({ isActive }) =>
        cn(
          "rounded-md bg-transparent px-2 py-0.5 text-xs font-medium text-muted-foreground transition-colors",
          isActive && "bg-foreground/10 text-foreground",
        )
      }
    >
      {children}
    </NavLink>
  );
}
