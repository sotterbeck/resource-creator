import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import {
  ExportRandomTexture,
  OpenTextureFiles,
} from "../../wailsjs/go/internal/App";
import { Variant, VariantList } from "@/components/variant-list";
import { Plus } from "lucide-react";

export default function Random() {
  const [material, setMaterial] = useState<string>("");
  const [variants, setVariants] = useState<Variant[]>([]);

  async function handleAddVariant() {
    let resp = await OpenTextureFiles("tile");
    if (!resp) {
      return;
    }

    const newVariants = resp.map((v) => ({
      name: v.name,
      path: v.path,
      imgData: v.imgData,
    }));

    setVariants((prev) =>
      [...prev, ...newVariants]
        .sort((a, b) => a.name.localeCompare(b.name))
        .filter(
          (variant, index, self) =>
            self.findIndex((v) => v.name === variant.name) === index,
        ),
    );
  }

  async function handleGenerate() {
    await ExportRandomTexture(
      variants.map((v) => v.path),
      material,
    );
  }

  return (
    <div className="flex h-full w-full flex-col gap-4 bg-background p-6">
      <div className="space-y-1">
        <Label htmlFor="material">Material</Label>
        <Input
          id="material"
          placeholder="minecraft:stone"
          onChange={(e) => setMaterial(e.target.value)}
        />
      </div>
      <div className="flex flex-1 flex-col items-start gap-2">
        <div className="flex w-full items-center justify-between">
          <span className="text-sm">Variants</span>
          <Button variant="secondary" size="icon" onClick={handleAddVariant}>
            <Plus />
          </Button>
        </div>
        <div className="w-full flex-1 overflow-y-auto">
          <VariantList
            variants={variants}
            setVariants={setVariants}
            className="h-full w-full"
          />
        </div>
        <Button
          variant="secondary"
          className="ml-auto mt-4"
          disabled={!material || variants.length === 0}
          onClick={handleGenerate}
        >
          Generate
        </Button>
      </div>
    </div>
  );
}
