import { cn } from "@/lib/utils";
import { twMerge } from "tailwind-merge";
import React from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { MoreVertical } from "lucide-react";

export interface Variant {
  name: string;
  path: string;
  imgData: string;
}

export interface VariantListProps {
  variants: Variant[];
  setVariants: React.Dispatch<React.SetStateAction<Variant[]>>;
  className?: string;
}

export function VariantList({
  variants,
  className,
  setVariants,
}: VariantListProps) {
  return (
    <div className={twMerge("w-full rounded-md border p-1", className)}>
      {variants.map((variant, index) => (
        <div key={index} className={cn("flex items-center gap-4 p-1")}>
          <img
            src={variant.imgData}
            alt={variant.name}
            className="pixelated h-12 w-12"
          />
          <div className="flex w-full items-center justify-between pr-3">
            <span className="text-sm tabular-nums">{variant.name}</span>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" size="icon">
                  <MoreVertical />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent>
                <DropdownMenuItem
                  onClick={() =>
                    setVariants((prev) => prev.filter((_, i) => i !== index))
                  }
                >
                  Delete
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </div>
      ))}
    </div>
  );
}
