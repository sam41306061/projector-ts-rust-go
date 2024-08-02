import getConfigPath, { Operation } from "../config";

test("Running my first test", function () {
  const config = getConfigPath({});
  expect(config.operation).toEqual(Operation.Print);
  expect(config.args).toEqual([]);
});
