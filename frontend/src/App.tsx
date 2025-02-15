import { Route, Routes } from "react-router";
import Pattern from "@/pages/pattern";
import Random from "@/pages/random";
import { Layout } from "@/layout";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Pattern />} />
        <Route path="/random" element={<Random />} />
      </Route>
    </Routes>
  );
}

export default App;
