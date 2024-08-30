import getOpts from "../../opts";
import getConfg, { Operation } from "./config";
import { Projector } from "./projector";

const opts = getOpts();
const config = getConfg(opts);
const proj = Projector.fromConfig(config);

if (config.operation === Operation.Print) {
  if (config.args.length === 0) {
    console.log(JSON.stringify(proj.getValue));
  } else {
    const value = proj.getValue(config.args[0]);
    if (value) {
      console.log(value);
    }
  }
}

if (config.operation === Operation.Add) {
  proj.setValue(config.args[0], config.args[1]);
  proj.saved();
  console.log("added");
}

if (config.operation === Operation.Remove) {
  proj.setValue(config.args[0], config.args[1]);
  proj.saved();
  console.log("added");
}
console.log(getOpts());
