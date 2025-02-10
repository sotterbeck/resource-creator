import { Button } from "@/components/ui/button";
import { ThemeProvider } from "@/components/theme-provider";
import { Separator } from "@/components/ui/separator";
import { useState } from "react";
import { Input } from "@/components/ui/input";
import { ExportPatternCTM, OpenTextureFile } from "../wailsjs/go/internal/App";
import { ResolutionSlider } from "@/components/resolution-slider";
import { TilePreview } from "@/components/tile-preview";

function Spinner() {
  return null;
}

function App() {
  const [textureData, setTextureData] = useState<string>("");

  const [material, setMaterial] = useState<string>("");
  const [patternResolution, setPatternResolution] = useState<{
    width: number;
    height: number;
  }>({ width: 0, height: 0 });
  const [tileResolution, setTileResolution] = useState<number>(0);

  async function handleImportTexture() {
    const resp = await OpenTextureFile();
    setTextureData(resp.imgData);
    setPatternResolution({ width: resp.width, height: resp.height });
    setTileResolution(Math.min(resp.width, resp.height));
  }

  async function handleExport() {
    await ExportPatternCTM(material, tileResolution);
  }

  return (
    <ThemeProvider>
      <div className="flex h-screen flex-col">
        <header className="draggable h-9 w-full"></header>
        <div className="flex justify-center pb-12 pt-5">
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
            <span className="text-sm">Material</span>
            <Input
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
    </ThemeProvider>
  );
}

export default App;
