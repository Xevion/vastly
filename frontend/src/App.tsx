import Offer from "@components/Offer";
import { Search } from "@wails/go/main/App";
import { api } from "@wails/go/models";
import { useEffect, useState } from "react";

function App() {
  const [state, setState] = useState<api.ScoredOffer[] | null>(null);

  async function invoke() {
    const offers = await Search();
    console.log({ offer: offers[0] });
    setState(offers);
  }

  useEffect(() => {
    if (state === null) invoke();
  });

  return (
    <div id="App">
      <div className="p-4" onClick={invoke}>
        {state?.map((offer) => (
          <Offer offer={offer} />
        ))}
      </div>
    </div>
  );
}

export default App;
