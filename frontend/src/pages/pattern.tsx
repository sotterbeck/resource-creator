import { useEffect, useState } from "react";
import {
  ExportPatternCTM,
  OpenTextureFile,
} from "../../wailsjs/go/internal/App";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { ResolutionSlider } from "@/components/resolution-slider";
import { Input } from "@/components/ui/input";
import { TilePreview } from "@/components/tile-preview";
import { Label } from "@/components/ui/label";
import { EventsEmit, EventsOn } from "../../wailsjs/runtime";
import { internal } from "../../wailsjs/go/models";
import TextureFile = internal.TextureFile;

export default function Pattern() {
  const [textureData, setTextureData] = useState<string>("");

  const [material, setMaterial] = useState<string>("");
  const [patternResolution, setPatternResolution] = useState<{
    width: number;
    height: number;
  }>({ width: 0, height: 0 });
  const [tileResolution, setTileResolution] = useState<number>(0);

  useEffect(() => {
    // @ts-ignore
    window.runtime.OnFileDrop(function (x, y, paths) {
      EventsEmit("pattern-file-drop", paths);
    }, true);

    EventsOn("pattern-file-drop-response", (data: TextureFile) => {
      setTextureData(data.imgData);
      setPatternResolution({ width: data.width, height: data.height });
      setTileResolution(Math.min(data.width, data.height));
    });
  });

  async function handleImportTexture() {
    const resp = await OpenTextureFile("pattern");
    if (!resp) {
      return;
    }
    setTextureData(resp.imgData);
    setPatternResolution({ width: resp.width, height: resp.height });
    setTileResolution(Math.min(resp.width, resp.height));
  }

  async function handleExport() {
    await ExportPatternCTM(material, tileResolution);
  }

  return (
    <div className="drop-target flex h-full flex-col">
      <div className="flex justify-center pb-12 pt-6">
        <Button onClick={handleImportTexture}>Import Texture</Button>
      </div>
      <Separator className="dark:bg-black" />
      <div className="relative flex-grow space-y-4 bg-background p-6">
        <ResolutionSlider
          value={tileResolution}
          onValueChange={(value) => setTileResolution(value)}
          max={Math.min(patternResolution.width, patternResolution.height)}
          min={4}
        />
        <div className="space-y-1">
          <Label htmlFor="material">Material</Label>
          <Input
            id="material"
            placeholder="minecraft:stone"
            disabled={!textureData}
            value={material}
            onChange={(e) => setMaterial(e.target.value)}
          />
        </div>
        <TilePreview
          src={textureData}
          tileResolution={tileResolution}
          patternResolution={patternResolution}
        />
        <Button
          onClick={handleExport}
          className="absolute bottom-6 right-6"
          variant="secondary"
          disabled={!textureData || !material}
        >
          Generate
        </Button>
      </div>
    </div>
  );
}
