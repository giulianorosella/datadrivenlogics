from openpyxl import load_workbook, Workbook  # type: ignore
from typing import List, Any, Dict
from vars.vars import var_numbers, con_numbers, path_ex, rows_numbers  # type: ignore
import pandas as pd  # type: ignore
from utils.theorem_checker import th_checker  # type: ignore
from formulas.formulas_generator import formulas_prov9, formulas_z3, formulas  # type: ignore
import csv  # type: ignore


def divide(list: List[Any], length: int) -> Any:
    division: List[Any] = []
    for i in range(0, len(list), length):
        division.append(list[i : i + length])
    return division


# function that divide a given list in sublists, this is needed if our formulas exceed the length of an excel sheet


def move_to_ex() -> Any:
    formule: List[Any] = formulas_prov9(con_numbers)
    solutions: List[Any] = th_checker(formulas_z3(con_numbers))
    formule_divise: List[Any] = divide(formule, rows_numbers)
    solutions_divise: List[Any] = divide(solutions, rows_numbers)
    for i in range(len(formule_divise)):
        writer = pd.ExcelWriter(
            path_ex
            + str(var_numbers)
            + "var"
            + str(con_numbers)
            + "con"
            + str(i)
            + "part.xlsx",
            engine="openpyxl",
        )
        wb = writer.book
        df = pd.DataFrame(
            {
                "Formulas": formule_divise[i],
                "CL Solutions": solutions_divise[i],
            }
        )

        df.to_excel(writer, index=False)
        wb.save(
            path_ex
            + str(var_numbers)
            + "var"
            + str(con_numbers)
            + "con"
            + str(i)
            + "part.xlsx"
        )


move_to_ex()
# scomment to generate .xlsx files
# main function that move formulas and solutions to the excel sheet;
# it separates into different excel if the number of formulas exceeds the length of the sheet


def move_to_csv() -> Any:
    formule: List[Any] = formulas_prov9(con_numbers)
    solutions: List[Any] = th_checker(formulas_z3(con_numbers))
    formule_divise: List[Any] = divide(formule, rows_numbers)
    solutions_divise: List[Any] = divide(solutions, rows_numbers)
    if rows_numbers == 0:
        sheet: Dict[Any, Any] = {"Formulas": formule, "Solutions": solutions}
        df = pd.DataFrame(sheet)
        df.to_csv(
            "prover9/sheet_excel/"
            + str(var_numbers)
            + "var"
            + str(con_numbers)
            + "con.csv"
        )
    else:
        for i in range(len(formule_divise)):
            sheet1: Dict[Any, Any] = {
                "Formulas": formule_divise[i],
                "Solutions": solutions_divise[i],
            }
            df = pd.DataFrame(sheet1)
            df.to_csv(
                "prover9/sheet_excel/"
                + str(var_numbers)
                + "var"
                + str(con_numbers)
                + "con"
                + str(i)
                + "part.csv",
                index=False,
            )


# move_to_csv()

# scomment to generate .csv files
