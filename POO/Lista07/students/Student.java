package Lista07.students;

import java.util.Arrays;
import java.util.stream.DoubleStream;

public class Student {
    private String name;
    private String entry;
    private String addres;
    private double[] note;
    private String situation;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEntry() {
        return entry;
    }

    public void setEntry(String entry) {
        this.entry = entry;
    }

    public String getAddres() {
        return addres;
    }

    public void setAddres(String addres) {
        this.addres = addres;
    }

    public double[] getNote() {
        return note;
    }

    public void setNote(double[] note) {
        this.note = note;
    }

    public String getSituation() {
        return situation;
    }

    public void setSituation(String situation) {
        this.situation = situation;
    }

    public Student(String name, String entry, String addres, double[] note){
        this.name = name;
        this.entry = entry;
        this.addres = addres;
        this.note = note;

        this.situation = isStudentApproved();
    }

    public  Student(String name, String entry, String addres){
        this.name = name;
        this.entry = entry;
        this.addres = addres;
        for (int i = 0; i < 2; i++) this.note[i] = 0.0;

        this.situation = isStudentApproved() ;
    }

    public String isStudentApproved(){
        Double avg = studentAVGNote();

        return avg >= 6 ? "Approved" : "Reproved";
    }

    public Double studentAVGNote(){
        return DoubleStream.of(this.note).average().orElse(0.0);
    }

    @Override
    public String toString() {
        return "Student{" +
                "name='" + name + '\'' +
                ", entry='" + entry + '\'' +
                ", addres='" + addres + '\'' +
                ", notes=" + Arrays.toString(note) +
                ", situation='" + situation + '\'' +
                '}';
    }
}
