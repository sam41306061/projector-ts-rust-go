import getConfigPath, { Operation } from "../config";

test("Running my first test", function () {
  const config = getConfigPath({});
  expect(config.operation).toEqual(Operation.Print);
  expect(config.args).toEqual([]);
});
test("print key", function () {
  const config = getConfigPath({
    args: ["foo"],
  });
  expect(config.operation).toEqual(Operation.Print);
  expect(config.args).toEqual(["foo"]);
});
test("add key", function () {
  const config = getConfigPath({
    args: ["add", "foo", "bar"],
  });
  expect(config.operation).toEqual(Operation.Add);
  expect(config.args).toEqual(["foo", "bar"]);
});

test("remove key", function () {
  const config = getConfigPath({
    args: ["remove", "foo"],
  });
  expect(config.operation).toEqual(Operation.Remove);
  expect(config.args).toEqual(["foo", "bar"]);
});
