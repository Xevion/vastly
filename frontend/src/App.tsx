import { Greet } from "@wails/go/main/App";
import { useEffect, useState } from "react";

function App() {
  const [state, setState] = useState<string>("");

  useEffect(() => {
    Greet("World").then((result) => {
      setState(result);
    });
  });

  return (
    <div id="App">
      <div className="p-4">{state}</div>
    </div>
  );
}

export default App;
